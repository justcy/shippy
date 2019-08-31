build:
	protoc -I. --go_out=plugins=grpc:/Users/justcy/Documents/Develop/52xiyou/go/src/github.com/justcy/shippy/consignment-service proto/consignment/consignment.proto
	GOOS=linux GOARCH=amd64 go build
	docker build -t consignment-service .
run:
	docker run -d -p 50051:50051 consignment-service
