FROM golang:latest

WORKDIR /go/src/mentor-server

COPY . .

RUN go get -v ./...

# Tests dependency
RUN go get gopkg.in/gavv/httpexpect.v2

RUN go test ./test

EXPOSE 8080

CMD ["mentor-server"]

