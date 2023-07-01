//go:build ent

package data

import (
	"context"

	"gin-layout/biz"
	"gin-layout/logger"
)

type userRepo struct {
	data *Data
	log  logger.LogInfoFormat
}

func NewUserRepo(data *Data, log logger.LogInfoFormat) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log,
	}
}

func (u *userRepo) CreateUser(ctx context.Context, user *biz.User) error {
	u.data.db.User.Query().Limit(2).Offset(10)

	_, err := u.data.db.User.Create().
		SetAge(user.Age).
		SetName(user.Name).
		SetPhone(user.Phone).
		SetPassword(user.Password).
		Save(ctx)
	return err
}
