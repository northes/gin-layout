package response

import (
	"fmt"
	"net/http"
)

type Code int64

var responseCode = struct {
	Unknown             Code
	Success             Code
	BadRequest          Code
	Unauthorized        Code
	Forbidden           Code
	NotFound            Code
	InternalServerError Code
}{
	Unknown:             0,                              // 未定义默认
	Success:             http.StatusOK,                  // 200
	BadRequest:          http.StatusBadRequest,          // 400
	Unauthorized:        http.StatusUnauthorized,        // 401
	Forbidden:           http.StatusForbidden,           // 403
	NotFound:            http.StatusNotFound,            // 404
	InternalServerError: http.StatusInternalServerError, // 500
}

var responseCodeMsg = map[Code]string{
	responseCode.Unknown:             "未定义的错误信息",
	responseCode.Success:             "成功",
	responseCode.BadRequest:          "无效请求",
	responseCode.Unauthorized:        "未经授权的访问",
	responseCode.Forbidden:           "禁止的操作",
	responseCode.NotFound:            "找不到资源",
	responseCode.InternalServerError: "服务器内部错误",
}

func (code Code) Msg() string {
	if msg, ok := responseCodeMsg[code]; ok {
		return msg
	}
	return fmt.Sprintf("[%d]%s", code, responseCodeMsg[responseCode.Unknown])
}
