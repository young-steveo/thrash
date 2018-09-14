package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/young-steveo/thrash/lexer"
)

func TestPut(t *testing.T) {
	source := lexer.FromBytes([]byte("Put the whole of your heart into my hands\n"))
	tokens := lexer.Scan(source)
	program := Put(tokens.Advance(), tokens)
	assert.Equal(t, "Assignment:\n    Left: Identifier: my hands\n    Right: Arithmetic:\n        Left: Identifier: the whole\n        Op: of\n        Right: Identifier: your heart", program.String())
}

func TestPutLiteral(t *testing.T) {
	source := lexer.FromBytes([]byte("Put 123 into The Clock"))
	tokens := lexer.Scan(source)
	program := Put(tokens.Advance(), tokens)
	assert.Equal(t, "Assignment:\n    Left: Identifier: The Clock\n    Right: Number: 123", program.String())
}
