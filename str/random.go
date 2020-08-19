package generatestr

import (
	"crypto/rand"
	"encoding/base64"
)

// generate a random string with a specific number of characters
// how it works: https://play.golang.org/p/nJKCbRSno7H
func randStr(digit uint32) string {
	b := make([]byte, digit)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return base64.RawURLEncoding.EncodeToString(b)
}
