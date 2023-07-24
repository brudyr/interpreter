package parser

import (
	"testing"

	"github.com/brudyr/go-interpreter/ast"
	"github.com/brudyr/go-interpreter/lexer"
)

func TestLetStatements(t *testing.T) {
  input := `
  let x = 5;
  let y = 10;
  let foobar = 15;
  `

  lexer := lexer.New(input)
  parser := New(lexer)

  program := parser.ParseProgram()

  if program == nil {
    t.Fatalf("ParseProgram() returned nil")
  }

  if len(program.Statements) != 3 {
    t.Fatalf("The program does not contain 3 statements. Got: %d", len(program.Statements))
  }

  testCases := []struct {
    expected string
  }{
    {"x"},
    {"y"},
    {"foobar"},
  }

  for index, testCase := range testCases {
    statement := program.Statements[index]
    if !testLetStatement(t, statement, testCase.expected) {
      return
    }
  }
}

func testLetStatement(t *testing.T, s ast.Statement, expectedName string) bool {

  if s.TokenLiteral() != "let" {
    t.Errorf("Expected 'let' TokenLiteral. Got: %q", s.TokenLiteral())
    return false
  }

  letStatement, ok := s.(*ast.LetStatement)
  if !ok {
    t.Errorf("Statement is not a LetStatement. Got: %T", s)
    return false
  }

  if letStatement.Name.Value != expectedName {
    t.Errorf("Expected name Value: %s for LetStatement. Got: %s", expectedName, letStatement.Name.Value)
  }

  if letStatement.Name.TokenLiteral() != expectedName {
    t.Errorf("Expected name TokenLiteral: %s for LetStatement. Got: %s", expectedName, letStatement.Name.TokenLiteral())
  }

  return true
}
