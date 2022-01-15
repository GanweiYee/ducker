FROM golang:latest
WORKDIR /go/src/app
ADD . /go/src/app
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go mod tidy
EXPOSE 8848
CMD go run main.go