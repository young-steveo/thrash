package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/young-steveo/thrash/lexer"
)

func TestConditionGreater(t *testing.T) {
	source := lexer.FromBytes([]byte("If Tommy is greater than you\n"))
	tokens := lexer.Scan(source)

	ifExpr := If(tokens.Advance(), tokens)
	assert.Equal(t, "If: true", ifExpr.String())

	// tommy := Identifier(tokens.Advance(), tokens)
	// program := Greater(tommy, tokens.Advance(), tokens)
	// assert.Equal(t, "Boolean: true", program.String())
}
