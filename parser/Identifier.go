package parser

import (
	"github.com/young-steveo/thrash/ast"
	"github.com/young-steveo/thrash/token"
)

// Identifier prefix parser
func Identifier(t *token.Token, l *token.List) ast.Expression {
	return &ast.Identifier{Token: t}
}
