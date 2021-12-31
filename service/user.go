package service

import (
	"errors"

	"github.com/northes/gin-layout/dao"
	"github.com/northes/gin-layout/dao/user"
	"github.com/northes/gin-layout/utils"
)

func Login(u *user.User) (token string, err error) {
	var us *user.User
	us, err = dao.Repo.User.GetByAccount(us.Account)
	if err != nil {
		return "", err
	}

	if utils.Encrypt(us.Password, us.Salt) != us.Password {
		return "", errors.New("密码错误")
	}

	return utils.GenToken(us.ID)
}

func CreateUser(user *user.User) error {
	return dao.Repo.User.Create(user)
}
