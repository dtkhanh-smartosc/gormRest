genpb:
	protoc --proto_path=proto --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=proto/pb --go_out=proto/pb proto/*.proto