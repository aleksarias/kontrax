build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)src/github.com/aleksarias/kontrax/organization-service \
	proto/organization/organization.proto
	docker build -t organization-service .
run:
	docker run -p 50051\:50051 -e MICRO_SERVER_ADDRESS=\:50051 -e MICRO_REGISTRY=mdns organization-service
