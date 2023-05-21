package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type LoggerItem struct {
	Request string
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		ctxErrs := c.Errors.String()
		if len(ctxErrs) != 0 {
			fmt.Println(ctxErrs)
		}
	}
}
