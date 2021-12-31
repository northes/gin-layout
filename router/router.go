package router

import (
	"github.com/gin-gonic/gin"
	. "github.com/northes/gin-layout/controller"
	"github.com/northes/gin-layout/response"
)

func Init() *gin.Engine {
	r := gin.New()
	_ = r.SetTrustedProxies(nil)

	r.GET("/ping", PingHandler)

	user := r.Group("/user")
	{
		// 创建用户
		user.POST("/", CreateUserHandler)
		// 修改用户信息
		user.POST("/:id", UpdateUserHandler)
		// 获取用户
		user.GET("/:id", GetUserHandler)
	}

	// 找不到路由
	r.NoRoute(func(c *gin.Context) {
		response.Error(c, response.CodeNotFound)
	})

	return r
}
