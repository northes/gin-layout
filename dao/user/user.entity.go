package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Account  string `json:"account" gorm:"uniqueIndex,comment:用户账号"`
	Password string `json:"password" gorm:"comment:密码"`
	Salt     string `json:"salt" gorm:"comment:盐"`
}
