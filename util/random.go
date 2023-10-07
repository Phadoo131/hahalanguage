package util

import (
	"math/rand"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"
const runee = "./,><?|}{[]}W'+=-_)()*&^%$#@!"

func RandomInt (min, max int) int{
	return min + rand.Intn(max - min + 1)
}

func RandomString (n int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomByte() byte{
	k := len(alphabet)
	if k == 0 {
		return 0
	}
	randomIndex := rand.Intn(k)
	return alphabet[randomIndex]
}