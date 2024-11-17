package shared

import (
	"math/rand"
	"strings"
	"time"
)

func ToCamelCase(s string) string {
	if len(s) == 0 {
		return ""
	}

	return strings.ToUpper(s[:1]) + s[1:]
}

func RandomString(length int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, length)
	for i := range b {
		b[i] = letterRunes[r.Intn(len(letterRunes))]
	}

	return string(b)
}
