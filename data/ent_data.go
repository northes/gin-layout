//go:build ent

package data

import (
	"context"

	"github.com/go-redis/redis"
	"go.uber.org/zap"

	"gin-layout/config"
	"gin-layout/data/ent"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

const Name = "ent"

type Data struct {
	db  *ent.Client
	rdb *redis.Client
}

func NewDB(cfg *config.AppConf) (*Data, func(), error) {

	var client *ent.Client
	var err error

	switch cfg.DB.Driver {
	case config.DBDriver.MySQL:
		client, err = ent.Open(string(config.DBDriver.MySQL), cfg.MySQL.GetDSN())
	default:
		client, err = ent.Open(string(config.DBDriver.SQLite3), cfg.SQLite.GetDSN())
	}

	if err != nil {
		return nil, nil, err
	}

	if err = client.Schema.Create(context.Background()); err != nil {
		return nil, nil, err
	}

	d := &Data{
		db: client,
	}

	zap.L().Info("database init success")

	return d, func() {
		zap.L().Info("Clean up db clients...")
		_ = d.db.Close()
	}, nil
}

func (d *Data) Close() {
	_ = d.db.Close()
}
