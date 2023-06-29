//go:build gorm

package data

import (
	"context"

	"gin-layout/biz"
)

type userRepo struct {
	data *Data
}

func NewUserRepo(data *Data) biz.UserRepo {
	return &userRepo{
		data: data,
	}
}

func (u *userRepo) CreateUser(ctx context.Context, user *biz.User) error {
	return u.data.db.Create(&user).Error
}
