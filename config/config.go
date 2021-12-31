package config

import (
	"github.com/spf13/viper"
)

var (
	App   app
	DB    database
	Redis redis
	Log   log
)

func Init(confPath string) error {
	viper.SetConfigFile(confPath)
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	var c *config
	err = viper.Unmarshal(&c)
	if err != nil {
		return err
	}

	App = c.App
	DB = c.DB
	Redis = c.Redis
	Log = c.Log

	return nil
}
