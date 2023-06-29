//go:build gorm

package data

import (
	"gin-layout/biz"
	"gin-layout/config"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const Name = "gorm"

type Data struct {
	db *gorm.DB
}

func Init() (*Data, func(), error) {
	var db *gorm.DB
	var err error

	switch config.Conf().DB.Driver {
	case config.DBDriver.MySQL:
		db, err = gorm.Open(mysql.Open(config.Conf().MySQL.GetDSN()), &gorm.Config{})
	default:
		db, err = gorm.Open(sqlite.Open(config.Conf().SQLite.GetDSN()), &gorm.Config{})
	}
	if err != nil {
		return nil, nil, err
	}

	err = db.AutoMigrate(&biz.User{})
	if err != nil {
		return nil, nil, err
	}

	d := &Data{
		db: db,
	}
	return d, func() {}, nil
}
