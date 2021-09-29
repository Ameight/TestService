PROJECTNAME = $(shell basename "$(PWD)")

PROTO_PATH = $(LOCAL_BIN)/bin
.PHONY: pbgen
pbgen:
	$(LOCAL_BIN) protoc -I api/test --go_out=plugins=grpc:pkg/api api/test/test-api.proto

.PHONY: test
test:
	ls $(LOCAL_BIN)