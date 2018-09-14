package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/young-steveo/thrash/lexer"
)

func TestMinus(t *testing.T) {
	source := lexer.FromBytes([]byte("Put Everything without nothing into my heart\n"))
	tokens := lexer.Scan(source)
	program := Put(tokens.Advance(), tokens)
	assert.Equal(t, "Assignment:\n    Left: Identifier: my heart\n    Right: Arithmetic:\n        Left: Identifier: Everything\n        Op: without\n        Right: Null: nothing", program.String())
}
