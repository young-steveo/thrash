package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/young-steveo/thrash/lexer"
)

func TestString(t *testing.T) {
	source := lexer.FromBytes([]byte(`"They give me cat scratch fever"`))
	tokens := lexer.Scan(source)
	program := String(tokens.Advance(), tokens)
	assert.Equal(t, `String: "They give me cat scratch fever"`, program.String())
}
