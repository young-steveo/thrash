package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/young-steveo/thrash/lexer"
)

func TestNull(t *testing.T) {
	source := lexer.FromBytes([]byte("nothing"))
	tokens := lexer.Scan(source)
	program := Null(tokens.Advance(), tokens)
	assert.Equal(t, "Null: nothing", program.String())
}
