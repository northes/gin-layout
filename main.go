package main

import (
	"fmt"
	"gin-layout/biz"
	"gin-layout/config"
	"gin-layout/data"
	"gin-layout/logger"
	"gin-layout/routers"
	"gin-layout/service"
)

var (
	Version string
)

func newApp() (func(), error) {
	d, cleanup, err := data.NewData()
	if err != nil {
		return nil, err
	}

	if err = logger.Init(); err != nil {
		return nil, err
	}

	// User
	userRepo := data.NewUserRepo(d)
	userUseCase := biz.NewUserUsecase(userRepo)
	service.NewUserService(userUseCase)

	return cleanup, nil
}

func main() {
	err := config.Init()
	if err != nil {
		panic(err)
	}

	cleanup, err := newApp()
	if err != nil {
		panic(err)
	}
	defer cleanup()

	fmt.Printf("version: %s\n", Version)

	r := routers.SetupRouter()
	_ = r.Run(config.Conf().GetPort())
}
