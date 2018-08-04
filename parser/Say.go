package parser

import (
	"fmt"

	"github.com/young-steveo/thrash/ast"
	"github.com/young-steveo/thrash/token"
)

// Say prefix parser
func Say(t *token.Token, l *token.List) ast.Expression {
	expr, err := parseExpression(l, 0)
	if err != nil {
		fmt.Print(`Failed parsing print expression.`)
		return nil
	}
	return &ast.Print{Expression: expr}
}
