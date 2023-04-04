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
