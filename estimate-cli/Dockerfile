FROM alpine:latest

RUN mkdir -p /app
WORKDIR /app

ADD estimate.json /app/estimate.json
ADD estimate-cli /app/estimate-cli

CMD ["./estimate-cli"]
