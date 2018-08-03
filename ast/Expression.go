package ast

import "fmt"

// Expression is an interface for all expressions
type Expression interface {
	fmt.Stringer
}
