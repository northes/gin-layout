package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context) {
	c.JSON(http.StatusOK, MessageSuccess())
}

func SuccessWithData(c *gin.Context, data any) {
	rsp := MessageSuccess()
	rsp.Data = data
	c.JSON(http.StatusOK, rsp)
}

func BadRequest(c *gin.Context) {
	code := responseCode.BadRequest
	c.JSON(http.StatusBadRequest, &Message{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	})
}

func BadRequestWithRsp(c *gin.Context, rsp *Message) {
	c.JSON(http.StatusBadRequest, rsp)
}

func Forbidden(c *gin.Context) {
	code := responseCode.Forbidden
	ForbiddenWithRsp(c, &Message{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	})
}

func ForbiddenWithRsp(c *gin.Context, rsp *Message) {
	c.JSON(http.StatusForbidden, rsp)
}

func InternalServerError(c *gin.Context) {
	InternalServerErrorWithRsp(c, nil)
}

func InternalServerErrorWithMsg(c *gin.Context, msg string) {
	InternalServerErrorWithRsp(c, &Message{
		Code: responseCode.InternalServerError,
		Msg:  msg,
		Data: nil,
	})
}

func InternalServerErrorWithRsp(c *gin.Context, rsp *Message) {
	c.JSON(http.StatusInternalServerError, rsp)
}
