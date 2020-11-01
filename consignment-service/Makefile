build:
	protoc -I. --go_out=plugins=grpc:/Users/zhouyixin/Desktop/shippy/consignment-service proto/consignment/consignment.proto
	docker build -t consignment-service .
	protoc --proto_path=. --go_out=. --micro_out=. proto/consignment/consignment.proto
	protoc --proto_path=. --go_out=. --micro_out=. proto/vessel/vessel.proto
run:
	docker run -p 50051:50051 -e MICRO_SERVER_ADDRESS=:50051 service