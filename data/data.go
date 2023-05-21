package data

import (
	"context"
	"fmt"
	"gin-layout/config"
	"gin-layout/data/ent"
	"github.com/go-redis/redis"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

type Data struct {
	db  *ent.Client
	rdb *redis.Client
}

func NewData() (*Data, func(), error) {

	var client *ent.Client
	var err error

	switch config.Conf().DB.Driver {
	case config.DBDriver.MySQL:
		client, err = ent.Open("mysql", "<user>:<pass>@tcp(<host>:<port>)/<database>?parseTime=True")
	default:
		client, err = ent.Open("sqlite3", "file:test.db?cache=shared&_fk=1")
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

	return d, func() {
		fmt.Println("cleanup")
		_ = d.db.Close()
	}, nil
}

func (d *Data) Close() {
	_ = d.db.Close()
}