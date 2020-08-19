package generatestr

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

// generates a hash value from the given arguments
// how it works: https://play.golang.org/p/waeR6hsLDxS
func hash(args ...string) string {
	str := strings.Join(args, "")
	h := sha256.New()
	h.Write([]byte(str))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}
