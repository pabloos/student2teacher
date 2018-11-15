FROM golang:stretch

RUN mkdir -p /go/src/app

WORKDIR /go/src/app

CMD GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o std2tch.linux.latest