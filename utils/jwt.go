package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

const TokenExpireDuration = time.Hour * 2

var secret = []byte("gin-layouts")

type Claims struct {
	uuid uint
	jwt.StandardClaims
}

func GenToken(uid uint) (string, error) {
	c := Claims{
		uuid: uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "gin-layout-user",                          // 签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(secret)
}

func ParseToken(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("")
}
