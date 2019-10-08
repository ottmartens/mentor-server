FROM golang:latest

WORKDIR /go/src/mentor-server

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

EXPOSE 8080

CMD ["mentor-server"]

