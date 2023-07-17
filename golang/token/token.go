package token

type TokenType string

const (
  ILLEGAL TokenType = "ILLEGAL"
  EOF = "EOF"
  ASSIGN = "="
  LPAREN = "("
  RPAREN = ")"
  LBRACE = "{"
  RBRACE = "}"
  PLUS = "+"
  MINUS = "-"
  COMMA = ","
  SEMICOLON = ";"
  FUNCTION = "FUNCTION"
  LET = "LET"
  IDENT = "IDENT"
  INT = "INT"
)

type Token struct {
  Type TokenType
  Value string
}

func New(tt TokenType, val string) Token {
  tok := Token {
    Type: tt,
    Value: val,
  }

  return tok
}
