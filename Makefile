proto:
	protoc --go_out=internal --go-grpc_out=internal proto/storage.proto

docs:
	swag init --output docs/client --generalInfo cmd/client/main.go