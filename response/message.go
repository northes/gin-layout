package response

type Message struct {
	Code Code   `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func MessageSuccess() Message {
	code := responseCode.Success
	return Message{
		Code: code,
		Msg:  code.Msg(),
	}
}
