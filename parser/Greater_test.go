package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/young-steveo/thrash/lexer"
)

func TestConditionGreater(t *testing.T) {
	source := lexer.FromBytes([]byte("If Tommy is greater than Janice\n"))
	tokens := lexer.Scan(source)

	ifExpr := If(tokens.Advance(), tokens)
	assert.Equal(t, "If: Comparison:\n        Left: Identifier: Tommy\n        Op: greater\n        Right: Identifier: Janice\n    ", ifExpr.String())
}
