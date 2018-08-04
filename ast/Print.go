package ast

import (
	"fmt"
)

// Print an expression
type Print struct {
	Expression Expression
}

func (p *Print) String() string {
	return fmt.Sprintf(`Print: %s`, p.Expression)
}
