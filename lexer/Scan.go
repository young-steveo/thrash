package lexer

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/young-steveo/thrash/token"
)

var mode Mode

// Scan some source and return a slice of tokens
func Scan(s *Source) *token.List {
	var tokens []*token.Token

	mode = Statement

	var firstWord *token.Token
	for s.Current <= s.Length {
		skip(s)
		s.Start = s.Current
		if s.IsAtEnd() {
			tokens = append(tokens, token.Create(token.EOF, []byte{0}, s.Line))
			break
		}
		c := s.Advance()

		if isDigit(c) {
			tokens = append(tokens, scanNumber(s))
			continue
		}

		switch c {
		case '"':
			tokens = append(tokens, scanString(s))
		case '\n':
			mode = Statement
			tokens = append(tokens, token.Create(token.Newline, []byte{c}, s.Line))
			s.Line++
			firstWord = nil
		case ',':
			tokens = append(tokens, token.Create(token.Comma, []byte{c}, s.Line))
		case '.':
		case '!':
		case '?':
			// ignore punctuation
		case '\'':
			if s.Peek() == 's' && s.PeekAt(1) == ' ' {
				lexeme := []byte{c, s.Advance()}
				tokens = append(tokens, token.Create(token.Is, lexeme, s.Line))
				if firstWord.Type == token.Identifier || firstWord.Type == token.Pronoun {
					mode = Poetic
				}
			}
		default:
			t := scanWord(s)
			tokens = append(tokens, t)
			switch t.Type {
			case token.Is:
				if firstWord.Type == token.Identifier || firstWord.Type == token.Pronoun {
					mode = Poetic
				}
			case token.Says:
				mode = PoeticString
			}
		}
		if firstWord == nil || firstWord.Type == token.Newline {
			firstWord = tokens[len(tokens)-1]
		}
	}

	return token.ListFromTokens(tokens)
}

// Skip whitespace and comments
func skip(s *Source) {
	for {
		c := s.Peek()
		switch c {
		case ' ':
			s.Advance()
		case '\t':
			s.Advance()
		case '(':
			skipComment(s)
		default:
			return
		}
	}
}

func skipComment(s *Source) {
	for s.Peek() != ')' && !s.IsAtEnd() {
		s.Advance()
	}
	if s.Peek() == ')' {
		s.Advance()
	}
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func isUppercase(c byte) bool {
	return c >= 'A' && c <= 'Z'
}

func isAlpha(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func isWhitespace(c byte) bool {
	return c == '\t' || c == '\v' || c == '\f' || c == '\r' || c == ' ' || c == 0x85 || c == 0xA0
}

func scanString(s *Source) *token.Token {
	for s.Peek() != '"' && !s.IsAtEnd() {
		if s.Peek() == '\n' {
			s.Line++
		}
		s.Advance()
	}
	if s.IsAtEnd() {
		return token.ErrorToken(`Unterminated string.`, s.Line)
	}

	// closing quotation mark
	s.Advance()

	lexeme := s.Data[s.Start+1 : s.Current-1]
	return token.Create(token.String, lexeme, s.Line)
}

func scanNumber(s *Source) *token.Token {
	for isDigit(s.Peek()) {
		s.Advance()
	}

	// Look for a fractional part.
	if s.Peek() == '.' && isDigit(s.PeekAt(1)) {
		// Consume the "."
		s.Advance()

		for isDigit(s.Peek()) {
			s.Advance()
		}
	}
	lexeme := s.Data[s.Start:s.Current]
	return token.Create(token.Number, lexeme, s.Line)
}

func scanWord(s *Source) *token.Token {
	lexeme := scanLexeme(s)

	switch mode {
	case Poetic:
		if typ, isLiteral := token.Literals[string(lexeme)]; isLiteral {
			mode = Statement
			return token.Create(typ, lexeme, s.Line)
		}
		return scanPoeticNumber(s, lexeme)
	case PoeticString:
		return scanPoeticString(s, lexeme)
	}

	if typ, isKeyword := token.Keywords[strings.ToLower(string(lexeme))]; isKeyword {
		return token.Create(typ, lexeme, s.Line)
	}

	if isUppercase(lexeme[0]) {
		if token.IsTake(lexeme) {
			lexeme = scanSuffix(s, lexeme, []byte(` it to the top`))
			return token.Create(token.Continue, lexeme, s.Line)
		}
		if token.IsGive(lexeme) {
			lexeme = scanSuffix(s, lexeme, []byte(` back`))
			return token.Create(token.Return, lexeme, s.Line)
		}
		if token.IsBreak(lexeme) {
			lexeme = scanSuffix(s, lexeme, []byte(` it down`))
			return token.Create(token.Break, lexeme, s.Line)
		}
		return scanProperIdentifier(s, lexeme)
	}

	if token.IsArticle(lexeme) {
		skip(s)
		lexeme = scanLexeme(s)
		return token.Create(token.Identifier, lexeme, s.Line)
	}

	return token.ErrorToken(fmt.Sprintf(`unknown lexeme '%s'`, lexeme), s.Line)
}

func scanLexeme(s *Source) []byte {
	var lexeme []byte
	for isAlpha(s.Peek()) {
		s.Advance()
	}

	lexeme = s.Data[s.Start:s.Current]

	// ignore single quotes unless it matches the pattern 's\W
	if s.Peek() == '\'' {
		if s.PeekAt(1) == 's' && isWhitespace(s.PeekAt(2)) {
			return lexeme
		}
		// ignore it.
		s.Advance()
		return scanLexeme(s)
	}
	return lexeme
}

func scanProperIdentifier(s *Source, lexeme []byte) *token.Token {
	skip(s)
	for isUppercase(s.Peek()) {
		lexeme = scanLexeme(s)
		skip(s)
	}
	return token.Create(token.Identifier, lexeme, s.Line)
}

func scanSuffix(s *Source, lexeme []byte, suffix []byte) []byte {
	count := len(suffix)
	next := s.PeekSlice(count)
	if bytes.Equal(next, suffix) {
		lexeme = append(lexeme, suffix...)
		for i := 0; i < count; i++ {
			s.Advance()
		}
	}
	return lexeme
}

func scanPoeticNumber(s *Source, lexeme []byte) *token.Token {
	var c byte
	firstPeriod := true

	if !s.IsAtEnd() {
		c = s.Advance()
	}
	for c != '\n' {
		if c == '(' {
			skipComment(s)
		}
		if isAlpha(c) || isWhitespace(c) {
			lexeme = append(lexeme, c)
		} else if firstPeriod && c == '.' {
			firstPeriod = false
			lexeme = append(lexeme, c)
		}
		if s.IsAtEnd() {
			break
		}
		c = s.Advance()
	}
	if c == '\n' {
		s.Retreat()
	}
	lexeme = bytes.Join(bytes.Fields(lexeme), []byte{' '})
	mode = Statement
	return token.Create(token.PoeticNumber, lexeme, s.Line)
}

func scanPoeticString(s *Source, lexeme []byte) *token.Token {
	for !s.IsAtEnd() {
		c := s.Advance()
		if c == '\n' {
			s.Retreat()
			break
		}
		lexeme = append(lexeme, c)
	}
	mode = Statement
	return token.Create(token.String, lexeme, s.Line)
}
