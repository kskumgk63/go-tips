package generatestr

import (
	"crypto/rand"
)

// add symbols as necessary
const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// generate a random string with a specific number of characters
// how it works: https://play.golang.org/p/GYm9miEs8Ag
func generateRand(digit uint32) string {
	b := make([]byte, digit)
	rand.Read(b)

	var result string
	for _, v := range b {
		result += string(letters[int(v)%len(letters)])
	}
	return result
}
