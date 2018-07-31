package debugger

import (
	"fmt"

	"github.com/young-steveo/thrash/token"
)

// PrintTokens will output tokens to STDIO
func PrintTokens(tokens []*token.Token) {
	line := 0
	for i, t := range tokens {
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
