FROM golang:latest

WORKDIR /go/src/mentor-server

COPY . .

RUN go get -u -v ./...


RUN go test ./test

EXPOSE 8080

CMD ["mentor-server"]

