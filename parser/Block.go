package parser

import (
	"fmt"

	"github.com/young-steveo/thrash/ast"
	"github.com/young-steveo/thrash/token"
)

// Block parser
func Block(t *token.Token, l *token.List) ast.Expression {
	block := ast.NewBlock()
	for !l.IsAtEnd() && (t.Type != token.Newline || (t.Type == token.Newline && l.Peek().Type != token.Newline)) {
		expr, err := parseStatement(l)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		if expr != nil {
			block.Expressions = append(block.Expressions, expr)
		}
	}
	return block
}
