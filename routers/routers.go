package routers

import (
	"net/http"

	"gin-layout/config"
	"gin-layout/di"
	"gin-layout/routers/middleware"
	"gin-layout/service"

	"github.com/gin-gonic/gin"
)

func SetupRouter(cfg *config.AppConf) (*gin.Engine, error) {
	gin.SetMode(cfg.Mode)

	r := gin.New()
	r.Use(middleware.Logger(), gin.Recovery())

	// home
	{
		r.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "hello")
		})
	}

	r.Use(middleware.JWT())

	// user
	{
		var userSvc *service.UserService
		if err := di.Invoke(func(us *service.UserService) { userSvc = us }); err != nil {
			return nil, err
		}
		r.POST("/user", userSvc.CreateUser)
	}

	return r, nil
}
