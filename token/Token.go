package token

// Token is a single token from the source
type Token struct {
	Type   Type
	Lexeme []byte
	Line   int
}

// Create does what it says
func Create(t Type, lexeme []byte, line int) *Token {
	return &Token{t, lexeme, line}
}

// ErrorToken creates an Error Token
func ErrorToken(msg string, line int) *Token {
	return Create(Error, []byte(msg), line)
}
