package dao

import (
	"fmt"

	"github.com/northes/gin-layout/config"

	"github.com/northes/gin-layout/dao/user"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Repo *repo

type repo struct {
	User *user.UserRepo
}

func Init() error {
	db, err := gorm.Open(mysql.Open(dsn()), &gorm.Config{})
	if err != nil {
		return err
	}

	Repo = &repo{
		User: user.NewUserRepo(db),
	}

	if err = db.AutoMigrate(&user.User{}); err != nil {
		return err
	}

	return nil
}

func dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB.UserName,
		config.DB.Password,
		config.DB.Host,
		config.DB.Port,
		config.DB.DBName,
	)
}
