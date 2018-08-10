package ast

import (
	"fmt"

	"github.com/young-steveo/thrash/token"
)

// Comparison of one expression into a variable
type Comparison struct {
	Left  Expression
	Right Expression
	Op    *token.Token
}

func (c *Comparison) String() string {
	level++
	result := fmt.Sprintf("Comparison:\n%sLeft: %s\n%sOp: %s\n%sRight: %s", indent(), c.Left, indent(), c.Op, indent(), c.Right)
	level--
	return result
}

// Source fulfils the Expression interface
func (c *Comparison) Source() []byte {
	return concatBytes(c.Left.Source(), c.Op.Lexeme, c.Right.Source())
}
