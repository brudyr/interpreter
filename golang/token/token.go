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
)

type Token struct {
  Type TokenType
  Value string
}
