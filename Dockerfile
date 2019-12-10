FROM golang:latest

WORKDIR /go/src/mentor-server

COPY . .

RUN go get -v ./...

RUN go test ./test -v

EXPOSE 8080

CMD ["mentor-server"]

