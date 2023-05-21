package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"strconv"
)

var shareConf *AppConf

type AppConf struct {
	*Site   `mapstructure:"site"`
	*DB     `mapstructure:"db"`
	*Redis  `mapstructure:"redis"`
	*Logger `mapstructure:"logger"`
}

type Site struct {
	Mode string `mapstructure:"mode"`
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

func Init() error {
	conf := &AppConf{}
	viper.SetConfigFile("./conf/config.yaml")
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		_ = viper.Unmarshal(&conf)
	})
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	if err = viper.Unmarshal(&conf); err != nil {
		return err
	}
	shareConf = conf
	return nil
}

func Conf() *AppConf {
	return shareConf
}

func (c *AppConf) GetPort() string {
	return fmt.Sprintf(":%s", strconv.Itoa(c.Site.Port))
}
