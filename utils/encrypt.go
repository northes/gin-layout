package utils

import (
	"crypto/sha1"
	"encoding/hex"
)

func Encrypt(str, salt string) string {
	s := sha1.New()
	s.Write([]byte(str))
	return hex.EncodeToString(s.Sum(secret))
}
