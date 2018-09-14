package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/young-steveo/thrash/lexer"
)

func TestOver(t *testing.T) {
	source := lexer.FromBytes([]byte("Put Everything over 0 into my heart\n"))
	tokens := lexer.Scan(source)
	program := Put(tokens.Advance(), tokens)
	assert.Equal(t, "Assignment:\n    Left: Identifier: my heart\n    Right: Arithmetic:\n        Left: Identifier: Everything\n        Op: over\n        Right: Number: 0", program.String())
}
