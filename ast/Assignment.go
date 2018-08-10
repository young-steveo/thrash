package ast

import (
	"fmt"

	"github.com/young-steveo/thrash/token"
)

// Assignment of one expression into a variable
type Assignment struct {
	Left  Expression
	Right Expression
	Op    *token.Token
}

func (a *Assignment) String() string {
	level++
	result := fmt.Sprintf("Assignment:\n%sLeft: %s\n%sRight: %s", indent(), a.Left, indent(), a.Right)
	level--
	return result
}

// Source fulfils the Expression interface
func (a *Assignment) Source() []byte {
	return concatBytes(a.Left.Source(), a.Op.Lexeme, a.Right.Source())
}
