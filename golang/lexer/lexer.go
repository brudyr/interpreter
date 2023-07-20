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

  l.skipWhitespace()

  switch l.curr {
    case '=':
      if l.peakChar() == '=' {
        currentChar := l.curr
        l.readChar()
        tok = token.New(token.EQ, string(currentChar) + string(l.curr))
      } else {
        tok = token.New(token.ASSIGN, "=")
      }
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
    case '<':
      tok = token.New(token.LT, "<")
    case '>':
      tok = token.New(token.GT, ">")
    case '!':
      if l.peakChar() == '=' {
        currentChar := l.curr
        l.readChar()
        tok = token.New(token.NOT_EQ, string(currentChar) + string(l.curr))
      } else {
        tok = token.New(token.BANG, "!")
      }
    case '/':
      tok = token.New(token.SLASH, "/")
    case '*':
      tok = token.New(token.ASTERISK, "*")
    case 0:
      tok = token.New(token.EOF, "EOF")
    default:
      if isLetter(l.curr) {
        tokenVal := l.readIdent()
        return token.New(token.IdentifyTokenType(tokenVal), tokenVal)
      } else if isNumber(l.curr) {
        tokenVal := l.readNumber()
        return token.New(token.INT, tokenVal)
      } else {
        tok = token.New(token.ILLEGAL, "")
      }
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

func (l *Lexer) readIdent() string {
  startPos := l.pos
  for isLetter(l.curr) {
    l.readChar()
  }
  return l.input[startPos:l.pos]
}

func (l *Lexer) readNumber() string {
  startPos := l.pos
  for isNumber(l.curr) {
    l.readChar()
  }
  return l.input[startPos:l.pos]
}

func isLetter(char byte) bool {
  return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
}

func isNumber(char byte) bool {
  return (char >= '0' && char <= '9')
}

func (l *Lexer) skipWhitespace() {
  for l.curr == ' ' || l.curr == '\t' || l.curr == '\n' || l.curr == '\r' {
      l.readChar()
  }
}

func (l *Lexer) peakChar() byte {
  if l.readPos >= len(l.input) {
    return 0;
  } else {
    return l.input[l.readPos]
  }
}
