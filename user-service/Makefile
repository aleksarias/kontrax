build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)src/github.com/aleksarias/kontrax/user-service proto/user/user.proto
	docker build -t us.gcr.io/kontrax/user-service:latest .
	gcloud docker -- push us.gcr.io/kontrax/user-service:latest
deploy:
	kubectl replace -f ./deployments/*.yml