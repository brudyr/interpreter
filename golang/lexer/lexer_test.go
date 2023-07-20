package lexer

import (
	"fmt"
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
        testCase.expectedType,
        currentToken.Type,
      )
    }
  }
}

func TestMultiCharTokens(t *testing.T) {
  input := `let someNum = 6;
  let biggerNum = 12;

  let add = func(a, b) {
    a + b
  }

  let sum = add(someNum, biggerNum);
  !-/*5;
  1 < 3 > 2;

  if (1 < 3) {
    return true;
  } else {
    return false;
  }


  5 == 5;
  6 != 5;
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
    { token.BANG, "!" },
    { token.MINUS, "-" },
    { token.SLASH, "/" },
    { token.ASTERISK, "*" },
    { token.INT, "5" },
    { token.SEMICOLON, ";" },
    { token.INT, "1" },
    { token.LT, "<" },
    { token.INT, "3" },
    { token.GT, ">" },
    { token.INT, "2" },
    { token.SEMICOLON, ";" },
    { token.IF, "if" },
    { token.LPAREN, "(" },
    { token.INT, "1" },
    { token.LT, "<" },
    { token.INT, "3" },
    { token.RPAREN, ")" },
    { token.LBRACE, "{" },
    { token.RETURN, "return" },
    { token.TRUE, "true" },
    { token.SEMICOLON, ";" },
    { token.RBRACE, "}" },
    { token.ELSE, "else" },
    { token.LBRACE, "{" },
    { token.RETURN, "return" },
    { token.FALSE, "false" },
    { token.SEMICOLON, ";" },
    { token.RBRACE, "}" },
    { token.INT, "5" },
    { token.EQ, "==" },
    { token.INT, "5" },
    { token.SEMICOLON, ";" },
    { token.INT, "6" },
    { token.NOT_EQ, "!=" },
    { token.INT, "5" },
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
        testCase.expectedType,
        currentToken.Type,
      )
    } else {
      fmt.Printf(
        "[testcase #%d] - Correct TokenType - Expected: '%s' | Got: '%s'\n",
        index,
        testCase.expectedType,
        currentToken.Type,
      )
      fmt.Println(currentToken.Value)
    }

  }
}
