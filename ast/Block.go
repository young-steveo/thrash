package ast

import (
	"strings"
)

// Block of expressions
type Block struct {
	Expressions []Expression
}

// NewBlock is a constructor
func NewBlock() *Block {
	return &Block{Expressions: make([]Expression, 0)}
}

func (b *Block) String() string {
	result := make([]string, 0)
	for _, expr := range b.Expressions {
		result = append(result, expr.String())
	}
	return strings.Join(result, "\n")
}
