package utils

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

var (
	jwtSecret = []byte("just_navigation_server")
)

type UserJWTClaims struct {
	jwt.RegisteredClaims
	UserID int64 `json:"user_id"`
}

func CreateAndSignJWT(userID int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &UserJWTClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: "just_navigation",
		},
		UserID: userID,
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
