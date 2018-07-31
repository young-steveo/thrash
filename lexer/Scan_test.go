package lexer

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/young-steveo/thrash/token"
)

func TestScanAssignPoeticNumber(t *testing.T) {
	source := FromBytes([]byte(`Tommy was to work on the docks.`))
	tokens := Scan(source)
	expect := []*token.Token{
		&token.Token{Type: token.Identifier, Lexeme: []byte(`Tommy`), Line: 1},
		&token.Token{Type: token.Is, Lexeme: []byte(`was`), Line: 1},
		&token.Token{Type: token.PoeticNumber, Lexeme: []byte(`to work on the docks.`), Line: 1},
		&token.Token{Type: token.EOF, Lexeme: []byte{byte(0)}, Line: 1},
	}
	for i, tok := range expect {
		assert.Equal(t, tok, tokens[i])
	}
}

func TestScanAssignNumber(t *testing.T) {
	source := FromBytes([]byte(`the girl is 16`))
	tokens := Scan(source)
	expect := []*token.Token{
		&token.Token{Type: token.Identifier, Lexeme: []byte(`the girl`), Line: 1},
		&token.Token{Type: token.Is, Lexeme: []byte(`is`), Line: 1},
		&token.Token{Type: token.Number, Lexeme: []byte(`16`), Line: 1},
		&token.Token{Type: token.EOF, Lexeme: []byte{byte(0)}, Line: 1},
	}
	for i, tok := range expect {
		assert.Equal(t, tok, tokens[i])
	}
}

func TestScanAssignLiteral(t *testing.T) {
	source := FromBytes([]byte(`My Heart is empty`))
	tokens := Scan(source)
	expect := []*token.Token{
		&token.Token{Type: token.Identifier, Lexeme: []byte(`the girl`), Line: 1},
		&token.Token{Type: token.Is, Lexeme: []byte(`is`), Line: 1},
		&token.Token{Type: token.Number, Lexeme: []byte(`16`), Line: 1},
		&token.Token{Type: token.EOF, Lexeme: []byte{byte(0)}, Line: 1},
	}
	for i, tok := range expect {
		assert.Equal(t, tok, tokens[i])
	}
}

func TestStringLiteral(t *testing.T) {
	source := FromBytes([]byte(`shout "Get on up!"`))
	tokens := Scan(source)
	expect := []*token.Token{
		&token.Token{Type: token.Say, Lexeme: []byte(`shout`), Line: 1},
		&token.Token{Type: token.String, Lexeme: []byte(`Get on up!`), Line: 1},
		&token.Token{Type: token.EOF, Lexeme: []byte{byte(0)}, Line: 1},
	}
	for i, tok := range expect {
		assert.Equal(t, tok, tokens[i])
	}
}

func TestNewLines(t *testing.T) {
	source := FromBytes([]byte("Gina is at the diner all day.\n\nShe is working for her man. She brings home her pay."))
	tokens := Scan(source)
	expect := []*token.Token{
		&token.Token{Type: token.Identifier, Lexeme: []byte(`Gina`), Line: 1},
		&token.Token{Type: token.Is, Lexeme: []byte(`is`), Line: 1},
		&token.Token{Type: token.PoeticNumber, Lexeme: []byte(`at the diner all day.`), Line: 1},
		&token.Token{Type: token.Newline, Lexeme: []byte{'\n'}, Line: 1},
		&token.Token{Type: token.Newline, Lexeme: []byte{'\n'}, Line: 2},
		&token.Token{Type: token.Pronoun, Lexeme: []byte(`She`), Line: 3},
		&token.Token{Type: token.Is, Lexeme: []byte(`is`), Line: 3},
		&token.Token{Type: token.PoeticNumber, Lexeme: []byte(`working for her man. She brings home her pay`), Line: 3},
		&token.Token{Type: token.EOF, Lexeme: []byte{byte(0)}, Line: 3},
	}
	for i, tok := range expect {
		assert.Equal(t, tok, tokens[i])
	}
}

func TestFunctions(t *testing.T) {
	source := FromBytes([]byte(
		`Midnight takes your heart and your soul
			Give back your heart
		
		My World is death
		The Fire was hot
		If Midnight taking My World, The Fire is My World
		    Shout The Fire`))
	tokens := Scan(source)
	expect := []*token.Token{
		&token.Token{Type: token.Identifier, Lexeme: []byte(`Midnight`), Line: 1},
		&token.Token{Type: token.Takes, Lexeme: []byte(`takes`), Line: 1},
		&token.Token{Type: token.Identifier, Lexeme: []byte(`your heart`), Line: 1},
		&token.Token{Type: token.And, Lexeme: []byte(`and`), Line: 1},
		&token.Token{Type: token.Identifier, Lexeme: []byte(`your soul`), Line: 1},
		&token.Token{Type: token.Newline, Lexeme: []byte{'\n'}, Line: 1},
		&token.Token{Type: token.Return, Lexeme: []byte(`Give back`), Line: 2},
		&token.Token{Type: token.Identifier, Lexeme: []byte(`your heart`), Line: 2},
		&token.Token{Type: token.Newline, Lexeme: []byte{'\n'}, Line: 2},
		&token.Token{Type: token.Newline, Lexeme: []byte{'\n'}, Line: 3},
		&token.Token{Type: token.Identifier, Lexeme: []byte(`My World`), Line: 4},
		&token.Token{Type: token.Is, Lexeme: []byte(`is`), Line: 4},
		&token.Token{Type: token.PoeticNumber, Lexeme: []byte(`death`), Line: 4},
		&token.Token{Type: token.Newline, Lexeme: []byte{'\n'}, Line: 4},
		&token.Token{Type: token.Identifier, Lexeme: []byte(`The Fire`), Line: 5},
		&token.Token{Type: token.Is, Lexeme: []byte(`was`), Line: 5},
		&token.Token{Type: token.PoeticNumber, Lexeme: []byte(`hot`), Line: 5},
		&token.Token{Type: token.Newline, Lexeme: []byte{'\n'}, Line: 5},
		&token.Token{Type: token.If, Lexeme: []byte(`If`), Line: 6},
		&token.Token{Type: token.Identifier, Lexeme: []byte(`Midnight`), Line: 6},
		&token.Token{Type: token.Taking, Lexeme: []byte(`taking`), Line: 6},
		&token.Token{Type: token.Identifier, Lexeme: []byte(`My World`), Line: 6},
		&token.Token{Type: token.Comma, Lexeme: []byte(`,`), Line: 6},
		&token.Token{Type: token.Identifier, Lexeme: []byte(`The Fire`), Line: 6},
		&token.Token{Type: token.Is, Lexeme: []byte(`is`), Line: 6},
		&token.Token{Type: token.Identifier, Lexeme: []byte(`My World`), Line: 6},
		&token.Token{Type: token.Newline, Lexeme: []byte{'\n'}, Line: 6},
		&token.Token{Type: token.Say, Lexeme: []byte(`Shout`), Line: 7},
		&token.Token{Type: token.Identifier, Lexeme: []byte(`The Fire`), Line: 7},
		&token.Token{Type: token.EOF, Lexeme: []byte{byte(0)}, Line: 7},
	}
	for i, tok := range expect {
		assert.Equal(t, tok, tokens[i])
	}
}
