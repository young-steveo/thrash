package ast

import (
	"fmt"

	"github.com/young-steveo/thrash/token"
)

// String literal
type String struct {
	Token *token.Token
}

func (s *String) String() string {
	return fmt.Sprintf(`String: "%s"`, s.Token)
}
