package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/young-steveo/thrash/lexer"
)

func TestPronoun(t *testing.T) {
	source := lexer.FromBytes([]byte("he"))
	tokens := lexer.Scan(source)
	program := Pronoun(tokens.Advance(), tokens)
	assert.Equal(t, "Pronoun: he", program.String())
}
