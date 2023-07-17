package lexer

import "github.com/brudyr/go-interpreter/token"

type Lexer struct {
  input string
  curr byte // the currently 'selected' char
  pos int // the positon of the current char
  readPos int // the reading position in the input (somewhere in the string after the current char)
}

func New(input string) *Lexer {
  l := &Lexer {input: input}
  l.readChar()
  return l
}

func (l *Lexer) NextToken() token.Token {
  var tok token.Token
  switch l.curr {
    case '=':
      tok = token.New(token.ASSIGN, "=")
    case '(':
      tok = token.New(token.LPAREN, "(")
    case ')':
      tok = token.New(token.RPAREN, ")")
    case '{':
      tok = token.New(token.LBRACE, "{")
    case '}':
      tok = token.New(token.RBRACE, "}")
    case '+':
      tok = token.New(token.PLUS, "+")
    case '-':
      tok = token.New(token.MINUS, "-")
    case ',':
      tok = token.New(token.COMMA, ",")
    case ';':
      tok = token.New(token.SEMICOLON, ";")
    case 0:
      tok = token.New(token.EOF, "EOF")
    default:
      tok = token.New(token.ILLEGAL, "")
  }
  l.readChar()
  return tok
}

func (l *Lexer) readChar() {
  if l.readPos >= len(l.input) {
    l.curr = 0
  } else {
    l.curr = l.input[l.readPos]
  }
  l.pos = l.readPos
  l.readPos++
}
