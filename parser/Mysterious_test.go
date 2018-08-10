package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/young-steveo/thrash/lexer"
)

func TestMysterious(t *testing.T) {
	source := lexer.FromBytes([]byte("mysterious"))
	tokens := lexer.Scan(source)
	program := Mysterious(tokens.Advance(), tokens)
	assert.Equal(t, "Mysterious", program.String())
}
