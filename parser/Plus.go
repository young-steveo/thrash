package parser

import (
	"fmt"

	"github.com/young-steveo/thrash/ast"
	"github.com/young-steveo/thrash/token"
)

// Plus infix parser
func Plus(left ast.Expression, t *token.Token, l *token.List) ast.Expression {
	right, err := parseExpression(l, TERM)
	if err != nil {
		fmt.Print(`Failed parsing right hand Plus expression.`)
		return nil
	}
	return &ast.Arithmetic{Left: left, Op: t, Right: right}
}
