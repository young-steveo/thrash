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
		assert.Equal(t, tok, tokens.Tokens[i])
	}
}

func TestScanComment(t *testing.T) {
	source := FromBytes([]byte(`(skip me) Tommy (foo) was (was he?) real (this is empty)`))
	tokens := Scan(source)
	expect := []*token.Token{
		&token.Token{Type: token.Identifier, Lexeme: []byte(`Tommy`), Line: 1},
		&token.Token{Type: token.Is, Lexeme: []byte(`was`), Line: 1},
		&token.Token{Type: token.PoeticNumber, Lexeme: []byte(`real`), Line: 1},
		&token.Token{Type: token.EOF, Lexeme: []byte{byte(0)}, Line: 1},
	}
	for i, tok := range expect {
		assert.Equal(t, tok, tokens.Tokens[i])
	}
}

func TestScanAssignPoeticString(t *testing.T) {
	source := FromBytes([]byte("Tommy says Give me a Sign!\nHe says you better do it"))
	tokens := Scan(source)
	expect := []*token.Token{
		&token.Token{Type: token.Identifier, Lexeme: []byte(`Tommy`), Line: 1},
		&token.Token{Type: token.Says, Lexeme: []byte(`says`), Line: 1},
		&token.Token{Type: token.String, Lexeme: []byte(`Give me a Sign!`), Line: 1},
		&token.Token{Type: token.Newline, Lexeme: []byte{'\n'}, Line: 1},
		&token.Token{Type: token.Pronoun, Lexeme: []byte(`He`), Line: 2},
		&token.Token{Type: token.Says, Lexeme: []byte(`says`), Line: 2},
		&token.Token{Type: token.String, Lexeme: []byte(`you better do it`), Line: 2},
		&token.Token{Type: token.EOF, Lexeme: []byte{byte(0)}, Line: 2},
	}
	for i, tok := range expect {
		assert.Equal(t, tok, tokens.Tokens[i])
	}
}

func TestScanAssignNumber(t *testing.T) {
	source := FromBytes([]byte(`the girl is 16.5`))
	tokens := Scan(source)
	expect := []*token.Token{
		&token.Token{Type: token.Identifier, Lexeme: []byte(`the girl`), Line: 1},
		&token.Token{Type: token.Is, Lexeme: []byte(`is`), Line: 1},
		&token.Token{Type: token.Number, Lexeme: []byte(`16.5`), Line: 1},
		&token.Token{Type: token.EOF, Lexeme: []byte{byte(0)}, Line: 1},
	}
	for i, tok := range expect {
		assert.Equal(t, tok, tokens.Tokens[i])
	}
}

func TestScanAssignPoeticLiteral(t *testing.T) {
	source := FromBytes([]byte(`My Heart is empty!`))
	tokens := Scan(source)
	expect := []*token.Token{
		&token.Token{Type: token.Identifier, Lexeme: []byte(`My Heart`), Line: 1},
		&token.Token{Type: token.Is, Lexeme: []byte(`is`), Line: 1},
		&token.Token{Type: token.Null, Lexeme: []byte(`empty`), Line: 1},
		&token.Token{Type: token.EOF, Lexeme: []byte{byte(0)}, Line: 1},
	}
	for i, tok := range expect {
		assert.Equal(t, tok, tokens.Tokens[i])
	}
}

func TestAssignment(t *testing.T) {
	source := FromBytes([]byte("Put true into my heart?\nPut my heart into your hands."))
	tokens := Scan(source)
	expect := []*token.Token{
		&token.Token{Type: token.Put, Lexeme: []byte(`Put`), Line: 1},
		&token.Token{Type: token.True, Lexeme: []byte(`true`), Line: 1},
		&token.Token{Type: token.Into, Lexeme: []byte(`into`), Line: 1},
		&token.Token{Type: token.Identifier, Lexeme: []byte(`my heart`), Line: 1},
		&token.Token{Type: token.Newline, Lexeme: []byte{'\n'}, Line: 1},
		&token.Token{Type: token.Put, Lexeme: []byte(`Put`), Line: 2},
		&token.Token{Type: token.Identifier, Lexeme: []byte(`my heart`), Line: 2},
		&token.Token{Type: token.Into, Lexeme: []byte(`into`), Line: 2},
		&token.Token{Type: token.Identifier, Lexeme: []byte(`your hands`), Line: 2},
		&token.Token{Type: token.EOF, Lexeme: []byte{byte(0)}, Line: 2},
	}
	for i, tok := range expect {
		assert.Equal(t, tok, tokens.Tokens[i])
	}
}

func TestStringLiteral(t *testing.T) {
	source := FromBytes([]byte("shout \"Get on\nup!\""))
	tokens := Scan(source)
	expect := []*token.Token{
		&token.Token{Type: token.Say, Lexeme: []byte(`shout`), Line: 1},
		&token.Token{Type: token.String, Lexeme: []byte("Get on\nup!"), Line: 2},
		&token.Token{Type: token.EOF, Lexeme: []byte{byte(0)}, Line: 2},
	}
	for i, tok := range expect {
		assert.Equal(t, tok, tokens.Tokens[i])
	}
}

func TestUnterminatedString(t *testing.T) {
	source := FromBytes([]byte(`shout "break me...`))
	tokens := Scan(source)
	expect := []*token.Token{
		&token.Token{Type: token.Say, Lexeme: []byte(`shout`), Line: 1},
		&token.Token{Type: token.Error, Lexeme: []byte("Unterminated string."), Line: 1},
		&token.Token{Type: token.EOF, Lexeme: []byte{byte(0)}, Line: 1},
	}
	for i, tok := range expect {
		assert.Equal(t, tok, tokens.Tokens[i])
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
		assert.Equal(t, tok, tokens.Tokens[i])
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
		&token.Token{Type: token.Equality, Lexeme: []byte(`is`), Line: 6},
		&token.Token{Type: token.Identifier, Lexeme: []byte(`My World`), Line: 6},
		&token.Token{Type: token.Newline, Lexeme: []byte{'\n'}, Line: 6},
		&token.Token{Type: token.Say, Lexeme: []byte(`Shout`), Line: 7},
		&token.Token{Type: token.Identifier, Lexeme: []byte(`The Fire`), Line: 7},
		&token.Token{Type: token.EOF, Lexeme: []byte{byte(0)}, Line: 7},
	}
	for i, tok := range expect {
		assert.Equal(t, tok, tokens.Tokens[i])
	}
}

func TestWhileContinue(t *testing.T) {
	source := FromBytes([]byte("While my heart is true\nTake it to the top"))
	tokens := Scan(source)
	expect := []*token.Token{
		&token.Token{Type: token.While, Lexeme: []byte(`While`), Line: 1},
		&token.Token{Type: token.Identifier, Lexeme: []byte(`my heart`), Line: 1},
		&token.Token{Type: token.Is, Lexeme: []byte(`is`), Line: 1},
		&token.Token{Type: token.True, Lexeme: []byte("true"), Line: 1},
		&token.Token{Type: token.Newline, Lexeme: []byte{'\n'}, Line: 1},
		&token.Token{Type: token.Continue, Lexeme: []byte("Take it to the top"), Line: 2},
		&token.Token{Type: token.EOF, Lexeme: []byte{byte(0)}, Line: 2},
	}
	for i, tok := range expect {
		assert.Equal(t, tok, tokens.Tokens[i])
	}
}

func TestUntilBreak(t *testing.T) {
	source := FromBytes([]byte("Until my heart is true\nBreak it down"))
	tokens := Scan(source)
	expect := []*token.Token{
		&token.Token{Type: token.While, Lexeme: []byte(`Until`), Line: 1},
		&token.Token{Type: token.Identifier, Lexeme: []byte(`my heart`), Line: 1},
		&token.Token{Type: token.Is, Lexeme: []byte(`is`), Line: 1},
		&token.Token{Type: token.True, Lexeme: []byte("true"), Line: 1},
		&token.Token{Type: token.Newline, Lexeme: []byte{'\n'}, Line: 1},
		&token.Token{Type: token.Break, Lexeme: []byte("Break it down"), Line: 2},
		&token.Token{Type: token.EOF, Lexeme: []byte{byte(0)}, Line: 2},
	}
	for i, tok := range expect {
		assert.Equal(t, tok, tokens.Tokens[i])
	}
}

func TestError(t *testing.T) {
	source := FromBytes([]byte(`foo`))
	tokens := Scan(source)
	expect := []*token.Token{
		&token.Token{Type: token.Error, Lexeme: []byte("unknown lexeme 'foo'"), Line: 1},
		&token.Token{Type: token.EOF, Lexeme: []byte{byte(0)}, Line: 1},
	}
	for i, tok := range expect {
		assert.Equal(t, tok, tokens.Tokens[i])
	}
}

func TestOwnershipAssignment(t *testing.T) {
	source := FromBytes([]byte(`Janie's got a gun.`))
	tokens := Scan(source)
	expect := []*token.Token{
		&token.Token{Type: token.Identifier, Lexeme: []byte(`Janie`), Line: 1},
		&token.Token{Type: token.Is, Lexeme: []byte(`'s`), Line: 1},
		&token.Token{Type: token.PoeticNumber, Lexeme: []byte(`got a gun.`), Line: 1},
	}
	for i, tok := range expect {
		assert.Equal(t, tok, tokens.Tokens[i])
	}
}

func TestIgnore(t *testing.T) {
	source := FromBytes([]byte(`Janie'saurus's got a gun.`))
	tokens := Scan(source)
	expect := []*token.Token{
		&token.Token{Type: token.Identifier, Lexeme: []byte(`Janie'saurus`), Line: 1},
		&token.Token{Type: token.Is, Lexeme: []byte(`'s`), Line: 1},
		&token.Token{Type: token.PoeticNumber, Lexeme: []byte(`got a gun.`), Line: 1},
	}
	for i, tok := range expect {
		assert.Equal(t, tok, tokens.Tokens[i])
	}
}

func TestGreaterCondition(t *testing.T) {
	source := FromBytes([]byte(`If Janie is greater than him`))
	tokens := Scan(source)
	expect := []*token.Token{
		&token.Token{Type: token.If, Lexeme: []byte(`If`), Line: 1},
		&token.Token{Type: token.Identifier, Lexeme: []byte(`Janie`), Line: 1},
		&token.Token{Type: token.Equality, Lexeme: []byte(`is`), Line: 1},
		&token.Token{Type: token.Greater, Lexeme: []byte(`greater`), Line: 1},
		&token.Token{Type: token.Than, Lexeme: []byte(`than`), Line: 1},
		&token.Token{Type: token.Pronoun, Lexeme: []byte(`him`), Line: 1},
		&token.Token{Type: token.EOF, Lexeme: []byte{byte(0)}, Line: 1},
	}
	for i, tok := range expect {
		assert.Equal(t, tok, tokens.Tokens[i])
	}
}

func TestScanListenTo(t *testing.T) {
	source := FromBytes([]byte(`Listen to your heart`))
	tokens := Scan(source)
	expect := []*token.Token{
		&token.Token{Type: token.Listen, Lexeme: []byte(`Listen to`), Line: 1},
		&token.Token{Type: token.Identifier, Lexeme: []byte(`your heart`), Line: 1},
		&token.Token{Type: token.EOF, Lexeme: []byte{byte(0)}, Line: 1},
	}
	for i, tok := range expect {
		assert.Equal(t, tok, tokens.Tokens[i])
	}
}
