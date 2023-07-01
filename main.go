package main

import (
	"gin-layout/config"
	"gin-layout/di"
	"gin-layout/routers"
)

var (
	Version string
)

func main() {
	_, err := di.BuildContainer()
	if err != nil {
		panic("[dig build container] " + err.Error())
	}

	var cfg *config.AppConf
	if err = di.Invoke(func(c *config.AppConf) { cfg = c }); err != nil {
		panic("[invoke config] " + err.Error())
	}

	r, err := routers.SetupRouter(cfg)
	if err != nil {
		panic("[setup router] " + err.Error())
	}
	_ = r.Run(cfg.Addr())
}
