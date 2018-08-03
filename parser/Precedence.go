package parser

// Precedence of operations
type Precedence int

// Precedence
const (
	DEFAULT Precedence = iota
	ASSIGNMENT
	CONDITIONAL
	SUM
	PRODUCT
	EXPONENT
	PREFIX
	POSTFIX
	CALL
)
