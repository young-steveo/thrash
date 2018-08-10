package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/young-steveo/thrash/lexer"
)

func TestBoolean(t *testing.T) {
	source := lexer.FromBytes([]byte("true"))
	tokens := lexer.Scan(source)
	program := Boolean(tokens.Advance(), tokens)
	assert.Equal(t, "Boolean: true", program.String())
}
