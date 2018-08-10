package ast

import (
	"fmt"

	"github.com/young-steveo/thrash/token"
)

// Null expression
type Null struct {
	Token *token.Token
}

func (n *Null) String() string {
	return fmt.Sprintf(`Null: %s`, n.Token)
}

// Source fulfils the Expression interface
func (n *Null) Source() []byte {
	return n.Token.Lexeme
}
