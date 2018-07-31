package lexer

// Mode for the lexer
type Mode int

// Modes
const (
	Statement Mode = iota // Normal Lexing
	Poetic                // Potential poetic literal or number

	PoeticNumber
	PoeticString
)
