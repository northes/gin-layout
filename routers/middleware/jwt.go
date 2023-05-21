package middleware

import (
	"fmt"
	"gin-layout/utils"
	"github.com/gin-gonic/gin"
	"strings"
)

func JWT() gin.HandlerFunc {

	return func(c *gin.Context) {
		authStr := c.GetHeader("Authorization")
		if len(authStr) == 0 {
			fmt.Println("auth str empty")
			c.Abort()
			return
		}

		jwtToken := strings.SplitN(authStr, " ", 2)
		if len(jwtToken) == 0 {
			fmt.Println("jwt token empty")
			c.Abort()
			return
		}

		claims, err := utils.ParseJwt(jwtToken[1])
		if err != nil {
			fmt.Println(err.Error())
			c.Abort()
			return
		}

		_ = claims
		c.Next()
	}
}
