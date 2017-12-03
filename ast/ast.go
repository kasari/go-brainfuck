package ast

import (
	"bytes"

	"github.com/kasari/go-brainfuck/token"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

type IncrementPointerStatement struct {
	Token token.Token
}

func (s *IncrementPointerStatement) statementNode()       {}
func (s *IncrementPointerStatement) TokenLiteral() string { return s.Token.Literal }
func (s *IncrementPointerStatement) String() string       { return s.Token.Literal }

type DecrementPointerStatement struct {
	Token token.Token
}

func (s *DecrementPointerStatement) statementNode()       {}
func (s *DecrementPointerStatement) TokenLiteral() string { return s.Token.Literal }
func (s *DecrementPointerStatement) String() string       { return s.Token.Literal }

type IncrementDataStatement struct {
	Token token.Token
}

func (s *IncrementDataStatement) statementNode()       {}
func (s *IncrementDataStatement) TokenLiteral() string { return s.Token.Literal }
func (s *IncrementDataStatement) String() string       { return s.Token.Literal }

type DecrementDataStatement struct {
	Token token.Token
}

func (s *DecrementDataStatement) statementNode()       {}
func (s *DecrementDataStatement) TokenLiteral() string { return s.Token.Literal }
func (s *DecrementDataStatement) String() string       { return s.Token.Literal }

type OutputStatement struct {
	Token token.Token
}

func (s *OutputStatement) statementNode()       {}
func (s *OutputStatement) TokenLiteral() string { return s.Token.Literal }
func (s *OutputStatement) String() string       { return s.Token.Literal }

type InputStatement struct {
	Token token.Token
}

func (s *InputStatement) statementNode()       {}
func (s *InputStatement) TokenLiteral() string { return s.Token.Literal }
func (s *InputStatement) String() string       { return s.Token.Literal }

type WhileStatement struct {
	Token      token.Token
	Statements []Statement
}

func (s *WhileStatement) statementNode()       {}
func (s *WhileStatement) TokenLiteral() string { return s.Token.Literal }
func (s *WhileStatement) String() string {
	var out bytes.Buffer

	out.WriteByte('[')
	for _, stmt := range s.Statements {
		out.WriteString(stmt.String())
	}
	out.WriteByte(']')

	return out.String()
}
