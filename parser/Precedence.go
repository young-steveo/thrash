package parser

// Precedence of operations
type Precedence int

// Precedence
const (
	DEFAULT Precedence = iota
	ASSIGNMENT
	OR
	AND
	EQUALITY
	COMPARISON
	TERM
	FACTOR
	UNARY
	POSTFIX
	CALL
)
