package parser

import (
	"fmt"

	"github.com/young-steveo/thrash/ast"
	"github.com/young-steveo/thrash/token"
)

// GreaterEqual infix parser
func GreaterEqual(left ast.Expression, t *token.Token, l *token.List) ast.Expression {
	err := l.Consume(token.As)
	if err != nil {
		fmt.Printf("Expected `as` after `%s` but saw `%s`\n", t.Lexeme, l.Peek().Lexeme)
		return nil
	}
	right, err := parseExpression(l, COMPARISON)
	if err != nil {
		fmt.Printf("Failed parsing right hand GreaterEqual expression: %s\n", err.Error())
		return nil
	}
	return &ast.Comparison{Left: left, Op: t, Right: right}
}
