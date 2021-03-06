mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
base_dir := $(notdir $(patsubst %/,%,$(dir $(mkfile_path))))

SERVICE ?= $(base_dir)

DOCKER_REGISTRY?=registry.uw.systems
DOCKER_REPOSITORY_NAMESPACE?=billing
DOCKER_ID?=billing
DOCKER_REPOSITORY_IMAGE=$(SERVICE)
DOCKER_REPOSITORY=$(DOCKER_REGISTRY)/$(DOCKER_REPOSITORY_NAMESPACE)/$(DOCKER_REPOSITORY_IMAGE)

K8S_NAMESPACE=$(DOCKER_REPOSITORY_NAMESPACE)
K8S_DEPLOYMENT_NAME=$(DOCKER_REPOSITORY_IMAGE)
K8S_CONTAINER_NAME=$(K8S_DEPLOYMENT_NAME)

BUILDENV :=
BUILDENV += CGO_ENABLED=0
GIT_HASH := $(CIRCLE_SHA1)
ifeq ($(GIT_HASH),)
  GIT_HASH := $(shell git rev-parse HEAD)
endif
LINKFLAGS :=-s -X main.gitHash=$(GIT_HASH) -extldflags "-static"
TESTFLAGS := -v -cover
EMPTY :=
SPACE := $(EMPTY) $(EMPTY)
join-with = $(subst $(SPACE),$1,$(strip $2))
LINTER := golangci-lint


.PHONY: install
install:
	GO111MODULE=on GOPRIVATE="github.com/utilitywarehouse/*" go mod download

$(LINTER):
	@ [ -e ./bin/$(LINTER) ] || wget -O - -q https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s latest

.PHONY: lint
lint: $(LINTER)
	./bin/$(LINTER) run $(LINT_FLAGS)

.PHONY: clean
clean:
	rm -f $(SERVICE)
	rm -f client

# builds our binary
$(SERVICE): clean
	$(BUILDENV) go build -o $(SERVICE) -a -ldflags '$(LINKFLAGS)' ./cmd/$(SERVICE)
	$(BUILDENV) go build -o client -a -ldflags '$(LINKFLAGS)' ./cmd/client

build: $(SERVICE)

.PHONY: test
test:
	$(BUILDENV) go test $(TESTFLAGS) ./...

.PHONY: all
all: clean $(LINTER) lint test build

docker-image:
	docker build -t $(DOCKER_REPOSITORY):local . --build-arg SERVICE=$(SERVICE) --build-arg GITHUB_TOKEN=$(GITHUB_TOKEN)

ci-docker-auth:
	@echo "Logging in to $(DOCKER_REGISTRY) as $(DOCKER_ID)"
	@docker login -u $(DOCKER_ID) -p $(UW_DOCKER_PASS) $(DOCKER_REGISTRY)

ci-docker-build: ci-docker-auth
	docker build -t $(DOCKER_REPOSITORY):$(CIRCLE_SHA1) . --build-arg SERVICE=$(SERVICE) --build-arg GITHUB_TOKEN=$(GITHUB_TOKEN)
	docker tag $(DOCKER_REPOSITORY):$(CIRCLE_SHA1) $(DOCKER_REPOSITORY):latest
	docker push $(DOCKER_REPOSITORY)

.PHONY: protos
protos:

UW_GIT :=github.com/utilitywarehouse
APP := kafka-topic-applier
PROTO_MAPPINGS := ""
GENERATED_PROTOS_DIR := ./internal/pb
PROTO_DIR := ./proto

protos:
	go get github.com/gogo/protobuf/protoc-gen-gogoslick
	go get google.golang.org/grpc

	rm -rf $(GENERATED_PROTOS_DIR)

	mkdir -pv $(GENERATED_PROTOS_DIR)/kta

	protoc \
  	    -I $(PROTO_DIR) \
	    -I .:$(GOPATH)/src \
            -I ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	    --gogoslick_out=$(PROTO_MAPPINGS),plugins=grpc:$(GENERATED_PROTOS_DIR)/kta $(PROTO_DIR)/kta.proto

	go generate ./...

mocks:
	go generate ./...
