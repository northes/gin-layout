package biz

import (
	"context"
	"time"
)

type User struct {
	ID       int64
	CreateAt time.Time
	UpdateAt time.Time
	Age      int64
	Name     string
	Phone    int64
	Password string
}

type UserRepo interface {
	CreateUser(ctx context.Context, user *User) error
}

type UserUsecase struct {
	repo UserRepo
}

func NewUserUsecase(repo UserRepo) *UserUsecase {
	return &UserUsecase{
		repo: repo,
	}
}

func (u *UserUsecase) Create(ctx context.Context, user *User) error {
	return u.repo.CreateUser(ctx, user)
}
