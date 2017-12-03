package lexer

import "github.com/kasari/go-brainfuck/token"

type Lexer struct {
	input string
	pos   int
	ch    byte
}

func New(input string) *Lexer {
	return &Lexer{input: input}
}

func (l *Lexer) readChar() byte {
	if l.pos >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.pos]
	}
	l.pos++
	return l.ch
}

func (l *Lexer) NextToken() token.Token {
	l.readChar()

	switch l.ch {
	case '+':
		return newToken(token.PLUS, l.ch)
	case '-':
		return newToken(token.MINUS, l.ch)
	case '>':
		return newToken(token.GT, l.ch)
	case '<':
		return newToken(token.LT, l.ch)
	case '[':
		return newToken(token.LBRACKET, l.ch)
	case ']':
		return newToken(token.RBRACKET, l.ch)
	case '.':
		return newToken(token.DOT, l.ch)
	case ',':
		return newToken(token.COMMA, l.ch)
	case 0:
		return newToken(token.EOF, 0)
	default:
		return newToken(token.IGNORE, l.ch)
	}
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(ch),
	}
}
