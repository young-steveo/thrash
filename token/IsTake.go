package token

import "bytes"

var take = []byte(`Take`)

// IsTake checks for keyword break
func IsTake(word []byte) bool {
	return bytes.Equal(word, take)
}
