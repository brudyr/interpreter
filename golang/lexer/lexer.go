package lexer

import "github.com/brudyr/go-interpreter/token"

type Lexer struct {
  input string
  curr byte
}

func New(input string) *Lexer {
  l := &Lexer {input: input}
  return l
}

func (l *Lexer) NextToken() token.Token {
  return token.Token {
    Type: token.COMMA,
    Value: ",",
  }
}

func (l *Lexer) readChar() {

}
