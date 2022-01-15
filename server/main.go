package main

import "github.com/GanweiYee/ducker/server/initialize"

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

func main() {
	initialize.Config()
	initialize.Redis()
	initialize.Xorm()

	Router := initialize.Router()

	Router.Run(":8848")
}
