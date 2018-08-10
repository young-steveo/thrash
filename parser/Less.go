package parser

import (
	"fmt"

	"github.com/young-steveo/thrash/ast"
	"github.com/young-steveo/thrash/token"
)

// Less infix parser
func Less(left ast.Expression, t *token.Token, l *token.List) ast.Expression {
	err := l.Consume(token.Than)
	if err != nil {
		fmt.Printf(`Expected "than" after "%s" but saw "%s"\n`, t.Lexeme, l.Peek().Lexeme)
		return nil
	}
	right, err := parseExpression(l, COMPARISON)
	if err != nil {
		fmt.Printf("Failed parsing right hand Less expression: %s\n", err.Error())
		return nil
	}
	return &ast.Comparison{Left: left, Op: t, Right: right}
}
