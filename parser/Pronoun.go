package parser

import (
	"github.com/young-steveo/thrash/ast"
	"github.com/young-steveo/thrash/token"
)

// Pronoun prefix parser
func Pronoun(t *token.Token, l *token.List) ast.Expression {
	return &ast.Pronoun{Token: t}
}
