package parser

import (
	"github.com/kasari/go-brainfuck/ast"
	"github.com/kasari/go-brainfuck/lexer"
	"github.com/kasari/go-brainfuck/token"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for !p.curTokenIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.PLUS:
		return &ast.IncrementDataStatement{Token: p.curToken}
	case token.MINUS:
		return &ast.DecrementDataStatement{Token: p.curToken}
	case token.GT:
		return &ast.IncrementPointerStatement{Token: p.curToken}
	case token.LT:
		return &ast.DecrementPointerStatement{Token: p.curToken}
	case token.DOT:
		return &ast.OutputStatement{Token: p.curToken}
	case token.COMMA:
		return &ast.InputStatement{Token: p.curToken}
	case token.LBRACKET:
		return p.parseWhileStatement()
	default:
		return nil
	}
}

func (p *Parser) parseWhileStatement() *ast.WhileStatement {
	stmt := &ast.WhileStatement{Token: p.curToken}
	stmt.Statements = []ast.Statement{}

	p.nextToken()

	for !p.curTokenIs(token.RBRACKET) && !p.curTokenIs(token.EOF) {
		s := p.parseStatement()
		if s != nil {
			stmt.Statements = append(stmt.Statements, s)
		}
		p.nextToken()
	}

	return stmt
}
