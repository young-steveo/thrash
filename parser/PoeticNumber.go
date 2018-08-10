package parser

import (
	"bytes"
	"fmt"
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
	value, err := strconv.ParseFloat(strings.Join(digits, ``), 64)
	if err != nil {
		fmt.Println(`Could not parse digits.`)
		return nil
	}
	return &ast.Number{Token: t, Value: value}
}
