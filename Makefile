gen:
	protoc --proto_path=infrastructure/grpc infrastructure/grpc/proto/*.proto --go_out=infrastructure/ --go-grpc_out=infrastructure/