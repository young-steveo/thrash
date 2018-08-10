package parser

import (
	"fmt"

	"github.com/young-steveo/thrash/ast"
	"github.com/young-steveo/thrash/token"
)

// Greater infix parser
func Greater(left ast.Expression, t *token.Token, l *token.List) ast.Expression {
	err := l.Consume(token.Than)
	if err != nil {
		fmt.Printf(`Expected "than" after "%s" but saw "%s"`, t.Lexeme, l.Peek().Lexeme)
		return nil
	}
	right, err := parseExpression(l, COMPARISON)
	if err != nil {
		fmt.Print(`Failed parsing right hand Greater expression.`)
		return nil
	}
	return &ast.Comparison{Left: left, Op: t, Right: right}
}
