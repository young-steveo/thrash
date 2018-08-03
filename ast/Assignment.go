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
	return fmt.Sprintf(`Assignment: %s %s %s`, a.Left, a.Op, a.Right)
}
