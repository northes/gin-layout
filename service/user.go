package service

import (
	"gin-layout/biz"
	"gin-layout/logger"
	"gin-layout/message"
	"gin-layout/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserService struct {
	user *biz.UserUsecase
	log  logger.LogInfoFormat
}

func NewUserService(user *biz.UserUsecase, log logger.LogInfoFormat) *UserService {
	return &UserService{
		user: user,
		log:  log,
	}
}

func (u *UserService) CreateUser(c *gin.Context) {
	req := new(message.UserCreateRequest)
	if err := c.ShouldBindJSON(req); err != nil {
		zap.L().Error("参数错误", zap.Error(err))
		response.BadRequest(c)
		return
	}
	err := u.user.Create(c, &biz.User{
		Age:      req.Age,
		Name:     req.Name,
		Phone:    req.Phone,
		Password: req.Password,
	})
	if err != nil {
		u.log.Error("用户创建失败", zap.Error(err))
		response.InternalServerErrorWithMsg(c, "用户创建失败")
		return
	}

	response.SuccessWithData(c, &message.UserCreateResponse{
		OK: true,
	})
}
