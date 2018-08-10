package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/young-steveo/thrash/lexer"
)

func TestConditionEquality(t *testing.T) {
	source := lexer.FromBytes([]byte("If Tommy is nobody\n"))
	tokens := lexer.Scan(source)

	ifExpr := If(tokens.Advance(), tokens)
	assert.Equal(t, "If: Comparison:\n        Left: Identifier: Tommy\n        Op: is\n        Right: Null: nobody\n    ", ifExpr.String())
}
