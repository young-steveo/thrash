package parser

import (
	"fmt"

	"github.com/young-steveo/thrash/ast"
	"github.com/young-steveo/thrash/token"
)

// Equality infix parser
func Equality(left ast.Expression, t *token.Token, l *token.List) ast.Expression {
	next := l.Peek()
	if next.Type == token.As {
		l.Advance()
		next = l.Peek()
		switch next.Type {
		case token.GreaterEqual:
			return GreaterEqual(left, l.Advance(), l)
		case token.LessEqual:
			return LessEqual(left, l.Advance(), l)
		default:
			fmt.Println("Failed parsing comparison, expected `less` or `great` after `as`")
			return nil
		}
	}
	switch next.Type {
	case token.Greater:
		return Greater(left, l.Advance(), l)
	case token.Less:
		return Less(left, l.Advance(), l)
	}
	right, err := parseExpression(l, EQUALITY)
	if err != nil {
		fmt.Printf("Failed parsing right hand Equality expression: %s\n", err.Error())
		return nil
	}
	return &ast.Comparison{Left: left, Op: t, Right: right}
}
