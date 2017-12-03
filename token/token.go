package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	PLUS     = "+"
	MINUS    = "-"
	GT       = ">"
	LT       = "<"
	LBRACKET = "["
	RBRACKET = "]"
	COMMA    = ","
	DOT      = "."

	IGNORE = "IGNORE"
	EOF    = "EOF"
)
