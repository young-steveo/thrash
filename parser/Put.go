package parser

import (
	"fmt"

	"github.com/young-steveo/thrash/ast"
	"github.com/young-steveo/thrash/token"
)

// Put prefix parser
func Put(t *token.Token, l *token.List) ast.Expression {
	right, err := parseExpression(l, 0)
	if err != nil {
		fmt.Print(`Failed parsing value for assignment expression`)
		return nil
	}

	l.Consume(token.Into)

	left, err := parseExpression(l, 0)
	if err != nil {
		fmt.Print(`Failed parsing identifier for assignment expression`)
		return nil
	}

	return &ast.Assignment{Left: left, Op: t, Right: right}
}
