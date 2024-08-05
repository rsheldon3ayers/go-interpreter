package parser

import (
	"github.com/rsheldon3ayers/go-interpreter/ast"
	"github.com/rsheldon3ayers/go-interpreter/lexer"
	"github.com/rsheldon3ayers/go-interpreter/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// Read two tokens, so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParsePRogram() *ast.Program {
	return nil
}
