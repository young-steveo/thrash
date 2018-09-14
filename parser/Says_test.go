package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/young-steveo/thrash/lexer"
)

func TestPoeticString(t *testing.T) {
	source := lexer.FromBytes([]byte(`Tommy says the work is all gone.`))
	tokens := lexer.Scan(source)
	tommy := Identifier(tokens.Advance(), tokens)
	program := Is(tommy, tokens.Advance(), tokens)
	assert.Equal(t, "Assignment:\n    Left: Identifier: Tommy\n    Right: String: \"the work is all gone.\"", program.String())
}
