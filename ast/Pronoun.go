package ast

import (
	"fmt"

	"github.com/young-steveo/thrash/token"
)

// Pronoun is a variable
type Pronoun struct {
	Token *token.Token
}

func (p *Pronoun) String() string {
	return fmt.Sprintf(`Pronoun: %s`, p.Token)
}

// Source fulfils the Expression interface
func (p *Pronoun) Source() []byte {
	return p.Token.Lexeme
}
