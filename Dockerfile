FROM alpine

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN update-ca-certificates

WORKDIR /app/
ADD ./app /app/
ADD ./.env /app/.env

ENTRYPOINT ["./app"]

# CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app