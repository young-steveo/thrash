package token

import "fmt"

// List of tokens for the parser
type List struct {
	Tokens  []*Token
	Start   int
	Current int
}

// ListFromTokens returns a list wrapper
func ListFromTokens(t []*Token) *List {
	return &List{Tokens: t, Start: 0, Current: 0}
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
