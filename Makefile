IMAGE="quay.io/utilitywarehouse/kafka-topic-applier"

.PHONY: install
install:
	go install \
			github.com/golang/mock/mockgen \
			github.com/bufbuild/buf/cmd/buf \
			google.golang.org/protobuf/cmd/protoc-gen-go \
			google.golang.org/grpc/cmd/protoc-gen-go-grpc

.PHONY: generate
generate: install
	rm -rf ./internal/pb
	buf generate
	go generate ./...

.PHONY: format
format:
	buf format proto -w

.PHONY: release
release:
	@sd "$(IMAGE):master" "$(IMAGE):$(VERSION)" $$(rg -l -- $(IMAGE) manifests/)
	@git add -- manifests/
	@git commit -m "Release $(VERSION)"
	@sd "$(IMAGE):$(VERSION)" "$(IMAGE):master" $$(rg -l -- "$(IMAGE)" manifests/)
	@git add -- manifests/
	@git commit -m "Clean up release $(VERSION)"
