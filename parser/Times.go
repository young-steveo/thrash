package parser

import (
	"fmt"

	"github.com/young-steveo/thrash/ast"
	"github.com/young-steveo/thrash/token"
)

// Times infix parser
func Times(left ast.Expression, t *token.Token, l *token.List) ast.Expression {
	right, err := parseExpression(l, FACTOR)
	if err != nil {
		fmt.Print(`Failed parsing right hand Times expression.`)
		return nil
	}
	return &ast.Arithmetic{Left: left, Op: t, Right: right}
}
