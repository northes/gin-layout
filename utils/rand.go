package utils

import (
	"math/rand"
	"time"
)

const (
	NumberLetters    = "0123456789"
	LowercaseLetters = "abcdefghijklmnopqrstuvwxyz"
	CapitalLetters   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	LongLetter       = NumberLetters + LowercaseLetters + CapitalLetters
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandKey(n int, str string) string {
	letter := []byte(str)
	if n <= 0 {
		return ""
	}
	b := make([]byte, n)
	arc := uint8(0)
	if _, err := rand.Read(b[:]); err != nil {
		return ""
	}
	for i, x := range b {
		arc = x & byte(len(letter)-1)
		b[i] = letter[arc]
	}
	return string(b)
}
