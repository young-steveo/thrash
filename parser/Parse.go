package parser

import (
	"fmt"

	"github.com/young-steveo/thrash/ast"
	"github.com/young-steveo/thrash/token"
)

// Parse a token list into an AST
func Parse(l *token.List) ast.Expression {
	program := ast.NewBlock()
	for !l.IsAtEnd() {
		expr, err := parseStatement(l)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		if expr != nil {
			program.Expressions = append(program.Expressions, expr)
		}
	}
	return program
}

func parseStatement(l *token.List) (ast.Expression, error) {
	expr, err := parseExpression(l, DEFAULT)
	if err != nil {
		return nil, err
	}
	l.Consume(token.Newline)
	return expr, nil
}

func parseExpression(l *token.List, pre Precedence) (ast.Expression, error) {
	t := l.Advance()
	switch t.Type {
	case token.EOF:
		return nil, nil
	case token.Error:
		return nil, fmt.Errorf(`%s`, string(t.Lexeme))
	}
	var left ast.Expression
	p, ok := prefixParsers[t.Type]
	if !ok {
		return nil, fmt.Errorf(`unknown prefix parser for token %s`, string(t.Lexeme))
	}
	left = p(t, l)

	if left == nil {
		return nil, fmt.Errorf(`could not parse prefix for "%s"`, string(t.Lexeme))
	}

	next := l.Peek()
	i, ok := infixParsers[next.Type]
	if ok {
		nextPre, ok := infixPrecedence[next.Type]
		for ok && pre < nextPre {
			t = l.Advance()
			left = i(left, t, l)
			next = l.Peek()
			nextPre, ok = infixPrecedence[next.Type]
		}
	}
	return left, nil
}
