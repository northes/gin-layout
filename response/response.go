package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type response struct {
	Code    Code
	Message string
	Data    interface{}
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &response{
		Code:    CodeSuccess,
		Message: CodeSuccess.GetMessage(),
		Data:    data,
	})
}

func Error(c *gin.Context, code Code) {
	c.JSON(http.StatusOK, &response{
		Code:    code,
		Message: code.GetMessage(),
		Data:    nil,
	})
}

func ErrorWithMessage(c *gin.Context, code Code, msg string) {
	c.JSON(http.StatusOK, &response{
		Code:    code,
		Message: msg,
		Data:    nil,
	})
}
