package parser

import (
	"github.com/young-steveo/thrash/ast"
	"github.com/young-steveo/thrash/token"
)

// Infix parser function
type Infix func(left ast.Expression, t *token.Token, l *token.List) ast.Expression
