package token

import "bytes"

var breakWord = []byte(`Break`)

// IsBreak checks for keyword break
func IsBreak(word []byte) bool {
	return bytes.Equal(word, breakWord)
}
