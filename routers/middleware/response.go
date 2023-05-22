package middleware

import (
	"gin-layout/errorx"
	"gin-layout/message"

	"github.com/gin-gonic/gin"
)

func Response() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// 统一错误处理
		err, ok := errorx.GetWithCtx(c)
		if ok {
			message.FailedWithMsg(c, err.ErrCode.Msg())
		}
	}
}
