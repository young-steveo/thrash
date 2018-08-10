package ast

import (
	"fmt"

	"github.com/young-steveo/thrash/token"
)

// Arithmetic of one expression into a variable
type Arithmetic struct {
	Left  Expression
	Right Expression
	Op    *token.Token
}

func (a *Arithmetic) String() string {
	level++
	result := fmt.Sprintf("Arithmetic:\n%sLeft: %s\n%sOp: %s\n%sRight: %s", indent(), a.Left, indent(), a.Op, indent(), a.Right)
	level--
	return result
}

// Source fulfils the Expression interface
func (a *Arithmetic) Source() []byte {
	return concatBytes(a.Left.Source(), a.Op.Lexeme, a.Right.Source())
}
