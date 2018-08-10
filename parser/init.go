package parser

import "github.com/young-steveo/thrash/token"

var prefixParsers map[token.Type]Prefix
var infixParsers map[token.Type]Infix
var infixPrecedence map[token.Type]Precedence

func init() {
	prefixParsers = make(map[token.Type]Prefix)
	prefixParsers[token.False] = Boolean
	prefixParsers[token.Identifier] = Identifier
	prefixParsers[token.Mysterious] = Mysterious
	prefixParsers[token.Null] = Null
	prefixParsers[token.Number] = Number
	prefixParsers[token.PoeticNumber] = PoeticNumber
	prefixParsers[token.Pronoun] = Pronoun
	prefixParsers[token.Put] = Put
	prefixParsers[token.Say] = Say
	prefixParsers[token.String] = String
	prefixParsers[token.True] = Boolean

	infixParsers = make(map[token.Type]Infix)
	infixParsers[token.Is] = Is
	infixParsers[token.Greater] = Greater
	infixParsers[token.Less] = Less
	infixParsers[token.Plus] = Plus
	infixParsers[token.Minus] = Minus
	infixParsers[token.Times] = Times
	infixParsers[token.Over] = Over

	infixPrecedence = make(map[token.Type]Precedence)
	infixPrecedence[token.Is] = ASSIGNMENT
	infixPrecedence[token.Equality] = EQUALITY
	infixPrecedence[token.Greater] = COMPARISON
	infixPrecedence[token.Less] = COMPARISON
	infixPrecedence[token.Plus] = TERM
	infixPrecedence[token.Minus] = TERM
	infixPrecedence[token.Times] = FACTOR
	infixPrecedence[token.Over] = FACTOR
}
