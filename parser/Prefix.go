package parser

import (
	"github.com/young-steveo/thrash/ast"
	"github.com/young-steveo/thrash/token"
)

// Prefix parser function
type Prefix func(t *token.Token, l *token.List) ast.Expression
