package ast

import (
	"fmt"

	"github.com/young-steveo/thrash/token"
)

// Boolean true or false!
type Boolean struct {
	Token *token.Token
}

func (b *Boolean) String() string {
	return fmt.Sprintf(`Boolean: %s`, b.Token)
}

// Source fulfils the Expression interface
func (b *Boolean) Source() []byte {
	return b.Token.Lexeme
}
