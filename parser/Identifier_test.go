package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/young-steveo/thrash/lexer"
)

func TestIdentifier(t *testing.T) {
	source := lexer.FromBytes([]byte("Tommy Two Toes"))
	tokens := lexer.Scan(source)
	program := Identifier(tokens.Advance(), tokens)
	assert.Equal(t, "Identifier: Tommy Two Toes", program.String())
}
