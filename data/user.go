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
	u.data.db.User.Query().Limit(2).Offset(10)

	_, err := u.data.db.User.Create().
		SetAge(user.Age).
		SetName(user.Name).
		SetPhone(user.Phone).
		SetPassword(user.Password).
		Save(ctx)
	return err
}
