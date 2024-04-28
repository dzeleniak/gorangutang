package parser

import (
	"gorangutang/ast"
	"gorangutang/lexer"
	"gorangutang/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curToken.Type != token.EOF {
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
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	//p.nextToken()
	//name := p.parseIdentifier()
	//p.nextToken()
	//if p.curToken.Type != token.EQ {
	//	panic("no equals sign")
	//}
	//p.nextToken()
	//value := p.parseExpression()
	//
	//return &ast.LetStatement{Name: name, Value: value}

	stmt := &ast.LetStatement{Token: p.curToken}

	if !p.peekTokenIs(token.IDENT) {
		return nil
	}
	p.nextToken()

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}
func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}
