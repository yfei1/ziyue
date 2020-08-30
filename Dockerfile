FROM golang:1.14-alpine
WORKDIR /go/src/app
COPY . .

RUN go build client.go
ENTRYPOINT ["./client"]
