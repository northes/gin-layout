package middleware

import (
	"bytes"
	"io"
	"net/http"
	"strings"
	"time"

	"gin-layout/errorx"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		timeStart := time.Now()
		raw := c.Request.URL.RawQuery
		path := c.Request.URL.Path
		method := c.Request.Method
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		if raw != "" {
			path = path + "?" + raw
		}

		if method == http.MethodPost {
			reqBody, _ := c.GetRawData()
			if len(reqBody) > 0 {
				c.Request.Body = io.NopCloser(bytes.NewBuffer(reqBody))
			}
			reqBodyStr := string(reqBody)
			reqBodyStr = strings.Replace(reqBodyStr, " ", "", -1)
			reqBodyStr = strings.Replace(reqBodyStr, "\n", "", -1)
			reqBodyStr = strings.Replace(reqBodyStr, "\r", "", -1)
			zap.L().Debug("request", zap.String("body", reqBodyStr))
		}

		c.Next()

		latency := time.Now().Sub(timeStart)

		errx, ok := errorx.GetWithCtx(c)
		if !ok {
			errx = errorx.SuccessErr.New()
		}

		zap.L().Info("[GIN]",
			zap.Int("status_code", statusCode),
			zap.Any("Latency", latency),
			zap.String("client_ip", clientIP),
			zap.String("method", method),
			zap.String("path", path),
			zap.Any("err_code", errx.ErrCode),
			zap.Errors("errors", errx.Errors),
		)
	}
}
