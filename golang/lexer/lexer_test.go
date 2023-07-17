package lexer

import (
	"testing"

	"github.com/brudyr/go-interpreter/token"
)

func TestSingleCharTokens(t *testing.T) {
  input := "(){}+-=,;"

  testCases := []struct {
    expectedType token.TokenType
    expectedVal string
  }{
    { token.LPAREN, "(" },
    { token.RPAREN, ")" },
    { token.LBRACE, "{" },
    { token.RBRACE, "}" },
    { token.PLUS, "+" },
    { token.MINUS, "-" },
    { token.ASSIGN, "=" },
    { token.COMMA, "," },
    { token.SEMICOLON, ";" },
    { token.EOF, "" },
  }
  
  l := New(input)

  for index, testCase := range testCases {

    currentToken := l.NextToken();

    if currentToken.Type != testCase.expectedType {
      t.Fatalf(
        "[testcase #%d] - Wrong TokenType - Expected: '%s' | Got: '%s'",
        index,
        currentToken.Type, testCase.expectedType,
      )
    }
  }
}

func TestMultiCharTokens(t *testing.T) {
  input := `let someNum = 6;
  let biggerNum = 12;

  let add = fn(a, b) {
    a + b
  }

  let sum = add(someNum, biggerNum);
  `

  testCases := []struct {
    expectedType token.TokenType
    expectedVal string
  }{
    { token.LET, "let" },
    { token.IDENT, "someNum" },
    { token.ASSIGN, "=" },
    { token.INT, "6" },
    { token.SEMICOLON, ";" },
    { token.LET, "let" },
    { token.IDENT, "biggerNum" },
    { token.ASSIGN, "=" },
    { token.INT, "12" },
    { token.SEMICOLON, ";" },
    { token.LET, "let" },
    { token.IDENT, "add" },
    { token.ASSIGN, "=" },
    { token.FUNCTION, "fn" },
    { token.LPAREN, "(" },
    { token.IDENT, "a" },
    { token.COMMA, "," },
    { token.IDENT, "b" },
    { token.RPAREN, ")" },
    { token.LBRACE, "{" },
    { token.IDENT, "a" },
    { token.PLUS, "+" },
    { token.IDENT, "b" },
    { token.RBRACE, "}" },
    { token.LET, "let" },
    { token.IDENT, "sum" },
    { token.ASSIGN, "=" },
    { token.IDENT, "add" },
    { token.LPAREN, "(" },
    { token.IDENT, "someNum" },
    { token.COMMA, "," },
    { token.IDENT, "biggerNum" },
    { token.RPAREN, ")" },
    { token.SEMICOLON, ";" },
    { token.EOF, "" },
  }

  l := New(input)

  for index, testCase := range testCases {

    currentToken := l.NextToken();

    if currentToken.Type != testCase.expectedType {
      t.Fatalf(
        "[testcase #%d] - Wrong TokenType - Expected: '%s' | Got: '%s'",
        index,
        currentToken.Type, testCase.expectedType,
      )
    }
  }
}
