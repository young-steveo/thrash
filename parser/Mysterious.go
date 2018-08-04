package parser

import (
	"github.com/young-steveo/thrash/ast"
	"github.com/young-steveo/thrash/token"
)

// Mysterious prefix parser
func Mysterious(t *token.Token, l *token.List) ast.Expression {
	return &ast.Mysterious{Token: t}
}
