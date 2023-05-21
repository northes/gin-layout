package routers

import (
	"gin-layout/config"
	"gin-layout/routers/middleware"
	"gin-layout/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter() *gin.Engine {
	gin.SetMode(config.Conf().Mode)

	r := gin.New()
	r.Use(gin.Recovery(), gin.Logger(), middleware.JWT(), middleware.Response())
	// 首页
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello")
	})

	r.POST("/user", service.User().CreateUser)

	return r
}
