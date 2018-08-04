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
