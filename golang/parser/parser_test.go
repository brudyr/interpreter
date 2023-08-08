package parser

import (
	"testing"

	"github.com/brudyr/go-interpreter/ast"
	"github.com/brudyr/go-interpreter/lexer"
)

func TestLetStatements(t *testing.T) {
  input := `
  let x 5;
  let = 10;
  let 12;
  `

  lexer := lexer.New(input)
  parser := New(lexer)

  program := parser.ParseProgram()
  checkParserErrors(t, parser)

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

func checkParserErrors(t *testing.T, p *Parser) {
  errors := p.GetErrors()
  if len(errors) == 0 {
    return
  }

  t.Errorf("Parser has %d errors", len(errors))
  for _, msg := range errors {
    t.Errorf("Parser error: %q", msg)
  }
  t.FailNow()
}
