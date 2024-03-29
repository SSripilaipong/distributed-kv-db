export PATH := $(PWD)/build:$(PATH)

PROTOBUF_FILES = \
	api/grpc/main.proto

gen: $(PROTOBUF_FILES) go-gen

$(PROTOBUF_FILES):
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative $@

go-gen: build-gen-tools
	go generate ./...

build-gen-tools: build-tstexcgen build-tstmckgen

build-tstexcgen:
	go build -o ./build/tstexcgen ./cmd/tstexcgen

build-tstmckgen:
	go build -o ./build/tstmckgen ./cmd/tstmckgen

.PHONY: gen $(PROTOBUF_FILES)
