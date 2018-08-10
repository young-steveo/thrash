package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/young-steveo/thrash/lexer"
)

func TestListen(t *testing.T) {
	source := lexer.FromBytes([]byte("Listen\n"))
	tokens := lexer.Scan(source)
	program := Listen(tokens.Advance(), tokens)
	assert.Equal(t, `Listen`, program.String())
}

func TestListenTo(t *testing.T) {
	source := lexer.FromBytes([]byte("Listen to your heart\n"))
	tokens := lexer.Scan(source)
	program := Listen(tokens.Advance(), tokens)
	assert.Equal(t, `Listen: Identifier: your heart`, program.String())
}
