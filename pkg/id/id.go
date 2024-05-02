package id

import (
	"fmt"
	"math/rand"
	"time"
)

const charset string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func RandomID() string {
	now := time.Now()
	return fmt.Sprintf("%s:%s", now.String(), generateRandomString(5))
}

func generateRandomString(size int) string {
	b := make([]byte, size)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
