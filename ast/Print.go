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

// Source fulfils the Expression interface
func (p *Print) Source() []byte {
	return concatBytes([]byte(`print`), []byte{}, p.Expression.Source())
}
