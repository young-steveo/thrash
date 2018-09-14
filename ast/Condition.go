package ast

import (
	"fmt"
)

// Condition of one expression into a variable
type Condition struct {
	Expression Expression
	Block      Expression
	Else       Expression
}

func (c *Condition) String() string {
	level++
	result := fmt.Sprintf("If: %s\n%s%s", c.Expression, indent(), c.Block)
	if c.Else != nil {
		result = fmt.Sprintf("%s\nElse: %s\n%s%s", result, c.Expression, indent(), c.Else)
	}
	level--
	return result
}

// Source fulfils the Expression interface
func (c *Condition) Source() []byte {
	if c.Else == nil {
		return concatBytes(c.Expression.Source(), []byte{}, c.Block.Source())
	}
	return concatBytes(c.Expression.Source(), c.Block.Source(), c.Else.Source())
}
