package ast

import "github.com/brudyr/go-interpreter/token"

type Node interface {
  TokenLiteral() string
}

type Statement interface {
  Node
  statementNode()
}

type Expression interface {
  Node
  expressionNode()
}

type Program struct {
  Statements []Statement
}

func (p *Program) TokenLiteral() string {
  if len(p.Statements) > 0 {
    return p.Statements[0].TokenLiteral()
  } else{
    return ""
  }
}

type Identifier struct {
  Token token.Token // this is always the IDENT token
  Value string
}

// implement expression interface
func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string { return i.Token.Value }

type LetStatement struct {
  Token token.Token // this is always the LET token
  Name *Identifier
  Value Expression
}

// Implement statement interface
func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Value }
