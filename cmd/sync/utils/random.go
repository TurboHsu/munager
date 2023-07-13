package utils

import (
	"fmt"
	"math/rand"
	"time"
)

// This function generates a random string with given length
func GenerateRandomString(len int) (ret string) {
	// Summon random string
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := 0; i < len; i++ {
		char := r.Intn(87)
		// Had better not to use backslash
		if char == 52 {
			char = -2
		}
		ret += fmt.Sprintf("%c", rune(char+40))
	}
	return
}
