package parser

import (
	"fmt"
	"strconv"

	"github.com/young-steveo/thrash/ast"
	"github.com/young-steveo/thrash/token"
)

// Number prefix parser
func Number(t *token.Token, l *token.List) ast.Expression {
	value, err := strconv.ParseFloat(string(t.Lexeme), 64)
	if err != nil {
		fmt.Println(err.Error())
	}
	return &ast.Number{Token: t, Value: value}
}
