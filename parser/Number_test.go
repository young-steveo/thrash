package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/young-steveo/thrash/lexer"
)

func TestNumber(t *testing.T) {
	source := lexer.FromBytes([]byte("8675309"))
	tokens := lexer.Scan(source)
	program := Number(tokens.Advance(), tokens)
	assert.Equal(t, "Number: 8675309", program.String())
}
