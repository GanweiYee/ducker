FROM golang:latest as builder
WORKDIR /go/src/app
ADD . /go/src/app
RUN go deploy -w GO111MODULE=on
RUN go deploy -w GOPROXY=https://goproxy.cn,direct
RUN go env -w CGO_ENABLED=0
RUN go mod tidy
EXPOSE 8848
CMD go build -o server .

FROM alpine:latest

