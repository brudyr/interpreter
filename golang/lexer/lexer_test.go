package lexer

import (
	"testing"

	"github.com/brudyr/go-interpreter/token"
)

func TestNextToken(t *testing.T) {
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
