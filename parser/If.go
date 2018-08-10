package parser

import (
	"fmt"

	"github.com/young-steveo/thrash/ast"
	"github.com/young-steveo/thrash/token"
)

var parsingCondition bool

// If prefix parser
func If(t *token.Token, l *token.List) ast.Expression {
	parsingCondition = true
	expr, err := parseExpression(l, DEFAULT)
	if err != nil {
		fmt.Print(`Failed parsing condition expression.`)
		return nil
	}

	err = l.Consume(token.Newline)
	if err != nil {
		fmt.Printf("Expected new line after if expression, but saw \"%s\"\n", l.Peek().Lexeme)
		return nil
	}
	parsingCondition = false
	block := Block(l.Peek(), l)

	// @todo else condition

	return &ast.Condition{Expression: expr, Block: block}
}
