package generatestr

import (
	"crypto/sha1"
	"fmt"
	"strings"
)

// generates a hash value from the given arguments
// how it works: https://play.golang.org/p/9yEaWqc3kPy
func generateHash(args ...string) string {
	str := strings.Join(args, "")
	h := sha1.New()
	h.Write([]byte(str))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}
