package routers

import (
	"net/http"

	"gin-layout/config"
	"gin-layout/routers/middleware"
	"gin-layout/service"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	gin.SetMode(config.Conf().Mode)

	r := gin.New()
	r.Use(middleware.Logger(), gin.Recovery())
	// 首页
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello")
	})

	r.Use(middleware.JWT())
	r.POST("/user", service.User().CreateUser)

	return r
}
