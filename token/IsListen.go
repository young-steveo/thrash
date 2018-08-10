package token

import "bytes"

var listenWord = []byte(`Listen`)

// IsListen checks for keyword break
func IsListen(word []byte) bool {
	return bytes.Equal(word, listenWord)
}
