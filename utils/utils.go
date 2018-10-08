package utils

import (
	"math/rand"
	"time"
)

const (
	letters = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	length  = len(letters)
)

func generateRandomString(l int) (string, error) {
	b := make([]byte, l)
	rand.Seed(time.Now().UnixNano())
	for i := range b {
		b[i] = byte(letters[rand.Intn(length)])
	}
	return string(b), nil
}

// func main() {
// 	for i := 0; i < 100; i++ {
// 		fmt.Println(generateRandomString(10))
// 	}
// }
