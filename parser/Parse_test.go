package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/young-steveo/thrash/lexer"
)

func TestParse(t *testing.T) {
	source := lexer.FromBytes([]byte(`Tommy was a cool dude`))
	tokens := lexer.Scan(source)
	program := Parse(tokens)
	assert.Equal(t, "Assignment:\n    Left: Identifier: Tommy\n    Right: Number: 144", program.String())
}
