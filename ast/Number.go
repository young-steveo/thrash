package ast

import (
	"fmt"

	"github.com/young-steveo/thrash/token"
)

// Number literal
type Number struct {
	Token *token.Token
	Value float64
}

func (n *Number) String() string {
	return fmt.Sprintf(`Number: %.10g`, n.Value)
}

// Source fulfils the Expression interface
func (n *Number) Source() []byte {
	return n.Token.Lexeme
}
