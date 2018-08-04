package ast

import (
	"fmt"

	"github.com/young-steveo/thrash/token"
)

// Mysterious expression
type Mysterious struct {
	Token *token.Token
}

func (m *Mysterious) String() string {
	return fmt.Sprint(`Mysterious`)
}
