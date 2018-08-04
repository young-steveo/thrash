package parser

import (
	"github.com/young-steveo/thrash/ast"
	"github.com/young-steveo/thrash/token"
)

// Number prefix parser
func Number(t *token.Token, l *token.List) ast.Expression {
	return &ast.Number{Token: t}
}
