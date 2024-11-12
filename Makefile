PROTO_DIR = api/proto
OUT_DIR = internal/interfaces/grpc
MODULE = github.com/gilsonyamada/user-api

.PHONY: proto
proto:
	protoc --proto_path=$(PROTO_DIR) \
	       --go_out=$(OUT_DIR) \
	       --go_opt=module=$(MODULE)/$(OUT_DIR) \
	       --go-grpc_out=$(OUT_DIR) \
	       --go-grpc_opt=module=$(MODULE)/$(OUT_DIR) \
	       $(PROTO_DIR)/*.proto