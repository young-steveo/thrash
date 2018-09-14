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

func TestConditionElse(t *testing.T) {
	source := lexer.FromBytes([]byte("If Tommy is nobody\nShout \"GotIt!\"\n\nElse\nShout My Life\n\n"))
	tokens := lexer.Scan(source)

	ifExpr := If(tokens.Advance(), tokens)
	assert.Equal(t, "If: Comparison:\n        Left: Identifier: Tommy\n        Op: is\n        Right: Null: nobody\n    Print: String: \"GotIt!\"\nPrint: Identifier: My Life", ifExpr.String())
}
