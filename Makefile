.PHONY: mocks
mocks:
	go generate ./...

.PHONY: docker-protoc
docker-protoc:
	rm -rf ./internal/pb
	docker run -v $$PWD:/defs namely/protoc-all -f ./proto/kta.proto -l go -o ./internal/pb/
