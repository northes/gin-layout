package config

import (
	"fmt"
	"strconv"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type AppConf struct {
	*Site   `mapstructure:"site"`
	*DB     `mapstructure:"db"`
	*Redis  `mapstructure:"redis"`
	*Logger `mapstructure:"logger"`
}

type Site struct {
	Mode string `mapstructure:"mode"`
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type Redis struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type Logger struct {
	Level      string `mapstructure:"level"`
	FileName   string `mapstructure:"file_name"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

func NewConfig() (*AppConf, error) {
	conf := &AppConf{}
	viper.SetConfigFile("./conf/config.yaml")
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		zap.L().Info("config changed", zap.Uint32("event", uint32(in.Op)))
		err := viper.Unmarshal(&conf)
		if err != nil {
			zap.L().Error("config changed unmarshal failed", zap.Error(err))
			return
		}
		zap.L().Info("config changed unmarshal success")
	})
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	if err = viper.Unmarshal(&conf); err != nil {
		return nil, err
	}
	return conf, nil
}

func (c *AppConf) Addr() string {
	return fmt.Sprintf("%s:%s", c.Site.Host, strconv.Itoa(c.Site.Port))
}
