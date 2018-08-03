package parser

import "github.com/young-steveo/thrash/token"

var prefixParsers map[token.Type]Prefix
var infixParsers map[token.Type]Infix
var infixPrecedence map[token.Type]Precedence

func init() {
	prefixParsers = make(map[token.Type]Prefix)
	prefixParsers[token.Identifier] = Identifier
	prefixParsers[token.True] = Boolean
	prefixParsers[token.False] = Boolean

	infixParsers = make(map[token.Type]Infix)
	infixParsers[token.Is] = Assignment

	infixPrecedence = make(map[token.Type]Precedence)
	infixPrecedence[token.Is] = ASSIGNMENT
}
