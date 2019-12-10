FROM golang:latest

WORKDIR /go/src/mentor-server

COPY . .

RUN go get ./...

# Tests dependency
RUN go get gopkg.in/gavv/httpexpect.v2

RUN go test ./test -v

EXPOSE 8080

CMD ["mentor-server"]

