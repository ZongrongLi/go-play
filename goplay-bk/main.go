package main

import (
	"github.com/tiancai110a/go-play/goplay-bk/conf"
	"github.com/tiancai110a/go-play/goplay-bk/server"
)

func main() {
	// 从配置文件读取配置
	conf.Init()

	// 装载路由
	r := server.NewRouter()
	r.Run(":3000")
}
