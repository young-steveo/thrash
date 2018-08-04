package parser

import (
	"bytes"
	"strconv"
	"strings"

	"github.com/young-steveo/thrash/ast"
	"github.com/young-steveo/thrash/token"
)

// PoeticNumber prefix parser
func PoeticNumber(t *token.Token, l *token.List) ast.Expression {
	words := bytes.Split(t.Lexeme, []byte{' '})
	length := len(words)
	index := length
	digits := make([]string, length)
	for index > 0 {
		index--
		digits[index] = strconv.Itoa(len(words[index]) % 10)
	}
	number := &token.Token{
		Type:   token.Number,
		Lexeme: []byte(strings.Join(digits, ``)),
		Line:   t.Line,
	}
	return &ast.Number{Token: number}
}
