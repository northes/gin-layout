package message

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type StatusCodeID int64

var StatusCode = struct {
	Success StatusCodeID
	Failed  StatusCodeID
}{
	Success: 200,
	Failed:  500,
}

type Response struct {
	Code StatusCodeID `json:"code"`
	Msg  string       `json:"msg"`
	Data any          `json:"data"`
}

var statusCodeMsg = map[StatusCodeID]string{
	StatusCode.Success: "success",
	StatusCode.Failed:  "failed",
}

func (s StatusCodeID) String() string {
	msg, ok := statusCodeMsg[s]
	if ok {
		return msg
	}
	return statusCodeMsg[StatusCode.Failed]
}

func Success(c *gin.Context) {
	SuccessWithData(c, nil)
}

func SuccessWithData(c *gin.Context, data any) {
	c.JSON(http.StatusOK, &Response{
		Code: StatusCode.Success,
		Msg:  StatusCode.Success.String(),
		Data: data,
	})
}

func Failed(c *gin.Context) {
	FailedWithMsg(c, StatusCode.Failed.String())
}

func FailedWithMsg(c *gin.Context, msg string) {
	c.JSON(http.StatusInternalServerError, &Response{
		Code: StatusCode.Failed,
		Msg:  msg,
		Data: nil,
	})
}
