package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/young-steveo/thrash/lexer"
)

func TestIs(t *testing.T) {
	source := lexer.FromBytes([]byte(`Tommy is a lean mean machine`))
	tokens := lexer.Scan(source)
	tommy := Identifier(tokens.Advance(), tokens)
	program := Is(tommy, tokens.Advance(), tokens)
	assert.Equal(t, "Assignment:\n    Left: Identifier: Tommy\n    Right: Number: 1447", program.String())
}
