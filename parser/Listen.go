package parser

import (
	"fmt"

	"github.com/young-steveo/thrash/ast"
	"github.com/young-steveo/thrash/token"
)

// Listen prefix parser
func Listen(t *token.Token, l *token.List) ast.Expression {
	next := l.Peek()
	switch next.Type {
	case token.Identifier:
		expr, err := parseExpression(l, 0)
		if err != nil {
			fmt.Println(`Failed parsing Listen expression.`)
			return nil
		}
		switch expr.(type) {
		case *ast.Identifier:
			return &ast.Listen{Var: expr}
		default:
			fmt.Println("`Listen to` expects an Identifier followed by a new line")
			return nil
		}
	case token.Error:
		fmt.Println(string(next.Lexeme))
		return nil
	}
	return &ast.Listen{}
}
