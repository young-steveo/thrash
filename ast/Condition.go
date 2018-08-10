package ast

import (
	"fmt"
)

// Condition of one expression into a variable
type Condition struct {
	Expression Expression
	Block      Expression
}

func (c *Condition) String() string {
	level++
	result := fmt.Sprintf("If: %s\n%s%s", c.Expression, indent(), c.Block)
	level--
	return result
}

// Source fulfils the Expression interface
func (c *Condition) Source() []byte {
	return concatBytes(c.Expression.Source(), []byte{}, c.Block.Source())
}
