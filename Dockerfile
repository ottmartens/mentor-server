FROM golang:latest

WORKDIR /go/src/mentor-server

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

RUN go test ./test

EXPOSE 8080

CMD ["mentor-server"]

