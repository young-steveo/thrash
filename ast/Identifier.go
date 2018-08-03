package ast

import (
	"fmt"

	"github.com/young-steveo/thrash/token"
)

// Identifier is a variable
type Identifier struct {
	Token *token.Token
}

func (i *Identifier) String() string {
	return fmt.Sprintf(`Identifier: %s`, i.Token)
}
