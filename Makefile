PROJECTNAME = $(shell basename "$(PWD)")

PROTO_PATH = $(LOCAL_BIN)/bin
.PHONY: pbgen
pbgen:
	$(PROTO_PATH) -I api/test/ --go_out=. api/test/test-api.proto 

#win /c/Users/egoro/go/proto/bin/protoc.exe -I api/test/ --go_out=. api/test/test-api.proto 