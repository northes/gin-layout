package errorx

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	ErrorCodeKey = "errMsg"
)

type Ierror interface {
	WithError(err error) *Error
	Error() string
}

type Error struct {
	ErrCode ErrorCodeID
	Errors  []error
}

type ErrorCodeID int64

const (
	UnknownErr ErrorCodeID = 0
	SuccessErr ErrorCodeID = 1
	ParamsErr  ErrorCodeID = 400
	ServerErr  ErrorCodeID = 500

	// User
	UserCreateErr                 ErrorCodeID = 1000
	UserNotExistOrPasswordFailErr ErrorCodeID = 1001
)

var errorCodeMsg = map[ErrorCodeID]string{
	UnknownErr: "未知错误",
	SuccessErr: "成功",
	ParamsErr:  "参数错误",
	ServerErr:  "服务器内部错误",

	// User
	UserCreateErr:                 "用户创建失败",
	UserNotExistOrPasswordFailErr: "用户不存在或密码错误",
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

func (e ErrorCodeID) New() *Error {
	return &Error{
		ErrCode: e,
	}
}

func (e ErrorCodeID) WithError(err error) *Error {
	return &Error{
		ErrCode: e,
		Errors:  []error{err},
	}
}

func (e ErrorCodeID) SetWithCtx(c *gin.Context, err error) {
	er := e.WithError(err)
	SetWithCtx(c, er)
}

func (e *Error) WithError(err error) *Error {
	e.Errors = append(e.Errors, err)
	return e
}

func (e *Error) Error() string {
	builder := strings.Builder{}
	builder.WriteString(e.ErrCode.Error())
	if len(e.Errors) > 0 {
		builder.WriteString(": ")
	}
	for _, err := range e.Errors {
		builder.WriteString("\n ")
		builder.WriteString(err.Error())
	}
	return builder.String()
}

func SetWithCtx(c *gin.Context, err Ierror) {
	if err == nil {
		return
	}
	if _, ok := err.(*Error); !ok {
		if errCode, ok := err.(ErrorCodeID); ok {
			err = errCode.New()
		} else {
			err = UnknownErr.WithError(err)
		}
	}

	c.Set(ErrorCodeKey, err)
}

func GetWithCtx(c *gin.Context) (*Error, bool) {
	code, ok := c.Get(ErrorCodeKey)
	if ok {
		err, ok := code.(*Error)
		return err, ok
	}

	return nil, false
}
