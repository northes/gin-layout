package response

type Code int64

const (
	CodeSuccess             Code = 200
	CodeNotFound                 = 404
	CodeBadParam                 = 400
	CodeInternalServerError      = 500
)

var MessageMap map[Code]string = map[Code]string{
	CodeSuccess:             "成功",
	CodeNotFound:            "无法找到数据",
	CodeBadParam:            "参数错误",
	CodeInternalServerError: "服务错误",
}

func (r Code) GetMessage() string {
	msg, ok := MessageMap[r]
	if !ok {
		return ""
	}
	return msg
}
