package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/young-steveo/thrash/lexer"
)

func TestSay(t *testing.T) {
	source := lexer.FromBytes([]byte(`Say "You really got me!"`))
	tokens := lexer.Scan(source)
	program := Say(tokens.Advance(), tokens)
	assert.Equal(t, `Print: String: "You really got me!"`, program.String())
}
