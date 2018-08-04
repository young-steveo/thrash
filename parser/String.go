package parser

import (
	"github.com/young-steveo/thrash/ast"
	"github.com/young-steveo/thrash/token"
)

// String prefix parser
func String(t *token.Token, l *token.List) ast.Expression {
	return &ast.String{Token: t}
}
