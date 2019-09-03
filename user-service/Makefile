build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/justcy/shippy/user-service proto/user/user.proto
	GOOS=linux GOARCH=amd64 go build
	docker build -t user-service .
run:
	docker run -p 50053:50051 \
         -e MICRO_SERVER_ADDRESS=:50051 \
         -e MICRO_REGISTRY=mdns \
         -e POSTGRES_HOST=database \
         -e POSTGRES_USER=postgres \
         -e POSTGRES_NAME=postgres \
         -e POSTGRES_PASSWORD=asdfweiiweo \
         user-service
