package main

import (
	"gin-layout/biz"
	"gin-layout/config"
	"gin-layout/data"
	"gin-layout/logger"
	"gin-layout/routers"
	"gin-layout/service"

	"go.uber.org/zap"
)

var (
	Version string
)

func newService(d *data.Data) error {
	// User
	userRepo := data.NewUserRepo(d)
	userUseCase := biz.NewUserUsecase(userRepo)
	service.NewUserService(userUseCase)

	return nil
}

func main() {
	// 配置初始化
	if err := config.Init(); err != nil {
		panic(err)
	}

	// 日志初始化
	if err := logger.Init(); err != nil {
		panic(err)
	}

	// 数据库初始化
	d, cleanup, err := data.Init()
	if err != nil {
		panic(err)
	}

	// 应用初始化
	if err = newService(d); err != nil {
		panic(err)
	}
	defer cleanup()

	zap.L().Info("server info:", zap.String("version", Version), zap.String("orm", data.Name), zap.String("gin-mode", config.Conf().Mode))

	r := routers.SetupRouter()
	_ = r.Run(config.Conf().Addr())
}
