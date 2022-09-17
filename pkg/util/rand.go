package util

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandInt(min, max int) int {
	return rand.Intn(max-min) + min
}

func RandInt64(min, max int64) int64 {
	return rand.Int63n(max-min) + min
}

func RandFloat64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func RandString(length int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func RandStringWithCharset(length int, charset string) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = rune(charset[rand.Intn(len(charset))])
	}
	return string(b)
}
