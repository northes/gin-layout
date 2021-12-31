package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/northes/gin-layout/dao/user"
	"github.com/northes/gin-layout/response"
	"github.com/northes/gin-layout/service"
)

func LoginHandler(c *gin.Context) {
	var u *user.User
	if err := c.ShouldBindJSON(&u); err != nil {
		response.Error(c, response.CodeBadParam)
		return
	}
	token, err := service.Login(u)
	if err != nil {
		response.Error(c, response.CodeInternalServerError)
		return
	}
	response.Success(c, gin.H{"token": token})
}

func CreateUserHandler(c *gin.Context) {
	var u *user.User
	if err := c.ShouldBindJSON(&u); err != nil {
		response.Error(c, response.CodeBadParam)
		return
	}

}

func GetUserHandler(c *gin.Context) {

}

func UpdateUserHandler(c *gin.Context) {

}
