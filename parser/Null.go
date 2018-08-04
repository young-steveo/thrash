package parser

import (
	"github.com/young-steveo/thrash/ast"
	"github.com/young-steveo/thrash/token"
)

// Null prefix parser
func Null(t *token.Token, l *token.List) ast.Expression {
	return &ast.Null{Token: t}
}
