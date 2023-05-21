package errorx

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
)

const (
	ErrorCodeKey = "errCode"
)

type ErrorCodeID int64

const (
	UnknownErr ErrorCodeID = 0
	ParamsErr  ErrorCodeID = 400

	// User
	UserCreateErr ErrorCodeID = 1000
)

var errorCodeMsg = map[ErrorCodeID]string{
	UnknownErr: "未知错误",
	ParamsErr:  "参数错误",

	// User
	UserCreateErr: "用户创建失败",
}

func (e ErrorCodeID) Error() string {
	return fmt.Sprintf("[%d]%s", e, e.Msg())
}

func (e ErrorCodeID) Msg() string {
	msg, ok := errorCodeMsg[e]
	if ok {
		return msg
	}
	return errorCodeMsg[UnknownErr]
}

func (e ErrorCodeID) SetInCtx(c *gin.Context, err error) {
	c.Set(ErrorCodeKey, e)
	if err != nil {
		_ = c.Error(errors.Join(e, err))
	}
}

func GetErrorCodeWithCtx(c *gin.Context) (ErrorCodeID, bool) {
	code, ok := c.Get(ErrorCodeKey)
	if ok {
		return code.(ErrorCodeID), true
	}
	return 0, false
}
