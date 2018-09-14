package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/young-steveo/thrash/lexer"
)

func TestPlus(t *testing.T) {
	source := lexer.FromBytes([]byte("Put Everything plus 1 into my hands\n"))
	tokens := lexer.Scan(source)
	program := Put(tokens.Advance(), tokens)
	assert.Equal(t, "Assignment:\n    Left: Identifier: my hands\n    Right: Arithmetic:\n        Left: Identifier: Everything\n        Op: plus\n        Right: Number: 1", program.String())
}
