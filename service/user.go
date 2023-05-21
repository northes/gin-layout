package service

import (
	"gin-layout/biz"
	"gin-layout/errorx"
	"gin-layout/message"
	"github.com/gin-gonic/gin"
)

type UserService struct {
	user *biz.UserUsecase
}

var shareUserService *UserService

func User() *UserService {
	return shareUserService
}

func NewUserService(user *biz.UserUsecase) {
	shareUserService = &UserService{
		user: user,
	}
}

func (u *UserService) CreateUser(c *gin.Context) {
	req := new(message.UserCreateRequest)
	if err := c.BindJSON(req); err != nil {
		errorx.ParamsErr.SetInCtx(c, err)
		return
	}
	err := u.user.Create(c, &biz.User{
		Age:      req.Age,
		Name:     req.Name,
		Phone:    req.Phone,
		Password: req.Password,
	})
	if err != nil {
		errorx.UserCreateErr.SetInCtx(c, err)
		return
	}

	message.SuccessWithData(c, &message.UserCreateResponse{
		OK: true,
	})
}
