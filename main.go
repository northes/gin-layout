package main

import (
	"flag"
	"fmt"

	"github.com/northes/gin-layout/config"
	"github.com/northes/gin-layout/dao"
	"github.com/northes/gin-layout/logger"
	"github.com/northes/gin-layout/router"
)

var confPath string

func init() {
	flag.StringVar(&confPath, "f", "./etc/config.yml", "config path, eg: -f ./etc/config.yml")
}

func main() {
	// 解析参数
	flag.Parse()
	// 初始化配置
	if err := config.Init(confPath); err != nil {
		panic(err)
	}
	// 初始化日志
	if err := logger.Init(); err != nil {
		panic(err)
	}
	// 初始化数据库
	if err := dao.Init(); err != nil {
		panic(err)
	}
	// 初始化路由
	r := router.Init()
	r.Run(fmt.Sprintf(":%s", config.App.Port))
}
