package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/young-steveo/thrash/lexer"
)

func TestBlock(t *testing.T) {
	source := lexer.FromBytes([]byte("\nTommy was a cool dude\n\n"))
	tokens := lexer.Scan(source)
	program := Block(tokens.Advance(), tokens)
	assert.Equal(t, "Assignment:\n    Left: Identifier: Tommy\n    Right: Number: 144", program.String())
}
