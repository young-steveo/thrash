package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/young-steveo/thrash/lexer"
)

func TestConditionGreaterEqual(t *testing.T) {
	source := lexer.FromBytes([]byte("If Tommy is as great as Janice\n"))
	tokens := lexer.Scan(source)

	ifExpr := If(tokens.Advance(), tokens)
	assert.Equal(t, "If: Comparison:\n        Left: Identifier: Tommy\n        Op: great\n        Right: Identifier: Janice\n    ", ifExpr.String())
}
