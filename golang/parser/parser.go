package parser

import (
	"github.com/brudyr/go-interpreter/ast"
	"github.com/brudyr/go-interpreter/lexer"
	"github.com/brudyr/go-interpreter/token"
)

type Parser struct {
  lexer *lexer.Lexer
  
  curToken token.Token
  peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
  p := &Parser{lexer: l}

  // initally fill curToken and peekToken
  p.nextToken()
  p.nextToken()

  return p
}

func (p *Parser) nextToken() {
  p.curToken = p.peekToken
  p.peekToken = p.lexer.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
  return nil
}
