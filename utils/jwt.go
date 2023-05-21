package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

var (
	jwtSecret = []byte("gin-layout")
)

func CreateAndSignJWT() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.RegisteredClaims{
		Issuer: "test",
	})
	tokenStr, err := token.SignedString(jwtSecret)
	return tokenStr, err
}

func ParseJwt(tokenStr string) (jwt.Claims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.Claims)
	if ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("jwt 解析失败")
}
