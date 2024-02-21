package utils

import (
	"math/rand"
)

func GenerateRefresh() string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	const refreshLen = 10

	const fn = "internal/utils/refreshToken/GenerateRefresh"
	b := make([]rune, refreshLen)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
