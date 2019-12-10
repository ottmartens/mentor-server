FROM golang:latest

WORKDIR /go/src/mentor-server

COPY . .

RUN go get -v ./...

# Tests dependency
RUN go get gopkg.in/gavv/httpexpect.v2

RUN cp /srv/mentor-server/.env ./test

RUN go test ./test -v

EXPOSE 8080

CMD ["mentor-server"]

