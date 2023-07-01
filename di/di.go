package di

import (
	"go.uber.org/dig"

	"gin-layout/biz"
	"gin-layout/config"
	"gin-layout/data"
	"gin-layout/logger"
	"gin-layout/service"
)

var container = dig.New()

func BuildContainer() (*dig.Container, error) {
	// config
	if err := container.Provide(config.NewConfig); err != nil {
		return nil, err
	}

	// logger
	if err := container.Provide(logger.NewLogger); err != nil {
		return nil, err
	}

	// DB
	if err := container.Provide(data.NewDB); err != nil {
		return nil, err
	}

	container.Provide(data.NewUserRepo)
	container.Provide(biz.NewUserUsecase)
	container.Provide(service.NewUserService)

	return container, nil
}

func Invoke(i any) error {
	return container.Invoke(i)
}
