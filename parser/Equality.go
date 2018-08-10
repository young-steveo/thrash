package parser

import (
	"fmt"

	"github.com/young-steveo/thrash/ast"
	"github.com/young-steveo/thrash/token"
)

// Equality infix parser
func Equality(left ast.Expression, t *token.Token, l *token.List) ast.Expression {
	next := l.Peek()
	switch next.Type {
	case token.Greater:
		return Greater(left, l.Advance(), l)
	case token.GreaterEqual:
	case token.Less:
	case token.LessEqual:
	}
	right, err := parseExpression(l, EQUALITY)
	if err != nil {
		fmt.Print(`Failed parsing right hand Equality expression.`)
		return nil
	}
	return &ast.Comparison{Left: left, Op: t, Right: right}
}
