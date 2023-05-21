package message

type UserCreateRequest struct {
	Age      int64  `json:"age"`
	Name     string `json:"name"`
	Phone    int64  `json:"phone"`
	Password string `json:"password"`
}

type UserCreateResponse struct {
	OK bool `json:"ok"`
}
