package token

import (
	"errors"
	"fmt"
)

// List of tokens for the parser
type List struct {
	Tokens  []*Token
	Start   int
	Current int
	Length  int
}

// ListFromTokens returns a list wrapper
func ListFromTokens(t []*Token) *List {
	return &List{Tokens: t, Start: 0, Current: 0, Length: len(t)}
}

// Print the list to STDOUT
func (l *List) Print() {
	line := 0
	for i, t := range l.Tokens {
		fmt.Printf(`%04d`, i)
		if i > 0 && t.Line == line {
			fmt.Print(`   | `)
		} else {
			line = t.Line
			fmt.Printf(`%4d `, line)
		}
		fmt.Printf("%02d %s\n", t.Type, t.Lexeme)
	}
}

// IsAtEnd checks if the tokens are consumed.
func (l *List) IsAtEnd() bool {
	return l.Current == l.Length
}

// Advance Current and return the previous Token
func (l *List) Advance() *Token {
	l.Current++
	return l.Tokens[l.Current-1]
}

// Peek at the current Token
func (l *List) Peek() *Token {
	return l.PeekAt(0)
}

// PeekAt a Token at position ahead of current
func (l *List) PeekAt(position int) *Token {
	offset := l.Current + position
	if offset >= l.Length {
		return nil
	}
	return l.Tokens[offset]
}

// Consume a token, or fail
func (l *List) Consume(typ Type) error {
	next := l.PeekAt(0)
	if next != nil && (next.Type == typ || next.Type == EOF) {
		l.Advance()
		return nil
	}

	// @todo make this explain what was expected with a lookup.
	return errors.New(`Expected a different token type`)
}
