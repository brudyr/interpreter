package token

type TokenType string

const (
  ILLEGAL TokenType = "ILLEGAL"
  EOF = "EOF"
  // single char
  ASSIGN = "="
  LPAREN = "("
  RPAREN = ")"
  LBRACE = "{"
  RBRACE = "}"
  PLUS = "+"
  MINUS = "-"
  COMMA = ","
  SEMICOLON = ";"
  BANG = "!"
  ASTERISK = "*"
  SLASH = "/"
  GT = ">"
  LT = "<"
  // double char
  EQ = "=="
  NOT_EQ = "!="
  // idents / keywords
  FUNCTION = "FUNCTION"
  LET = "LET"
  TRUE = "TRUE"
  FALSE = "FALSE"
  IF = "IF"
  ELSE = "ELSE"
  RETURN = "RETURN"
  IDENT = "IDENT"
  INT = "INT"
)

var keywords = map[string]TokenType {
  "func": FUNCTION,
  "let": LET,
  "true": TRUE,
  "false": FALSE,
  "if": IF,
  "else": ELSE,
  "return": RETURN,
}

func IdentifyTokenType(ident string) TokenType {
  if tok, ok := keywords[ident]; ok {
    return tok
  }
  return IDENT
}

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
