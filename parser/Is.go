package parser

import (
	"fmt"

	"github.com/young-steveo/thrash/ast"
	"github.com/young-steveo/thrash/token"
)

// Is infix parser
func Is(left ast.Expression, t *token.Token, l *token.List) ast.Expression {
	right, err := parseExpression(l, ASSIGNMENT)
	if err != nil {
		fmt.Print(`Failed parsing right hand assignment expression.`)
		return nil
	}
	return &ast.Assignment{Left: left, Op: t, Right: right}
}
