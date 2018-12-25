build:
	GOOS=linux GOARCH=amd64 go build
	docker build -t estimate-cli .

run:
	docker run -e MICRO_REGISTRY=mdns estimate-cli
