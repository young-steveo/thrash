package parser

import (
	"github.com/young-steveo/thrash/ast"
	"github.com/young-steveo/thrash/token"
)

// Boolean prefix parser
func Boolean(t *token.Token, l *token.List) ast.Expression {
	return &ast.Boolean{Token: t}
}
