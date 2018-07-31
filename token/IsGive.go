package token

import "bytes"

var give = []byte(`Give`)

// IsGive checks for keyword break
func IsGive(word []byte) bool {
	return bytes.Equal(word, give)
}
