package ast

import (
	"fmt"

	"github.com/young-steveo/thrash/token"
)

// Number literal
type Number struct {
	Token *token.Token
}

func (n *Number) String() string {
	return fmt.Sprintf(`Number: "%s"`, n.Token)
}
