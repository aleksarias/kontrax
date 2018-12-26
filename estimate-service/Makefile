build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)src/github.com/aleksarias/kontrax/estimate-service proto/estimate/estimate.proto
	docker build -t estimate-service .
run:
	docker run -p 50051\:50051 -e MICRO_SERVER_ADDRESS=\:50051 -e MICRO_REGISTRY=mdns estimate-service

