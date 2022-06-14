PROTO_OUT ?= pkg/pb

# Current dir
DIR ?= $(shell pwd)

# Dockerized protoc command
PROTOC_CMD ?= docker run --rm -v $(DIR):$(DIR) -w $(DIR) thethingsindustries/protoc

.PHONY: compile-proto-go
compile-proto-go: ## Compile protobuf contracts in order to generate Go stubs
	@rm -rf $(PROTO_OUT)
	@mkdir -p $(PROTO_OUT)
	@find proto -name "*.proto" -type f -exec $(PROTOC_CMD) -Iproto/ \
	     --go_opt=paths=source_relative \
	     --go_out=plugins=grpc:$(PROTO_OUT) {} \;
