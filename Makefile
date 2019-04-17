create_proto_definitions:
	protoc --go_out=. proto/*.proto
	protoc --go_out=plugins=grpc:. proto/*.proto