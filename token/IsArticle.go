package token

import "bytes"

var articles = [5][]byte{
	[]byte(`a`),
	[]byte(`an`),
	[]byte(`the`),
	[]byte(`my`),
	[]byte(`your`),
}

// IsArticle checks for article variable keywords
func IsArticle(word []byte) bool {
	for _, article := range articles {
		if bytes.Equal(article, word) {
			return true
		}
	}
	return false
}
