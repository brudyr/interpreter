package parser

import (
	"testing"

	"github.com/brudyr/go-interpreter/ast"
	"github.com/brudyr/go-interpreter/lexer"
)

func TestLetStatements(t *testing.T) {
	input := `
  let x =  5;
  let y = 10;
  let foobar = 12;
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

func TestReturnStatements(t *testing.T) {
	input := `
  return 12;
  return 42;
  return 1337;
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

	for _, statement := range program.Statements {
		returnStmt, ok := statement.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("Statement is not a ReturnStatement. Got: %T", statement)
			continue
		}

		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("Expected 'return' TokenLiteral. Got: %q", returnStmt.TokenLiteral())
		}
	}
}

func TestIdentifierExpression(t *testing.T) {
	input := "foobar;"

	lexer := lexer.New(input)
	parser := New(lexer)
	program := parser.ParseProgram()
	checkParserErrors(t, parser)

	if len(program.Statements) != 1 {
		t.Fatalf("program has unexpected number of statements. got: %d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("First program statement is not the expected type ast.ExpressionStatement instead it is: %T", program.Statements[0])
	}

	ident, ok := stmt.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("Exression is not of type *astIdentifier. It is: %T", stmt.Expression)
	}
	if ident.Value != "foobar" {
		t.Errorf("ident.Value ist not: %s got: %s", "foobar", ident.Value)
	}
	if ident.TokenLiteral() != "foobar" {
		t.Errorf("ident.TokenLiteral is not %s got: %s", "foobar", ident.TokenLiteral())
	}
}

func TestIntegerLiteralExpression(t *testing.T) {
	input := "5;"

  l := lexer.New(input)
  p := New(l)
  program := p.ParseProgram()
  checkParserErrors(t, p)

  if len(program.Statements) != 1 {
      t.Fatalf("program has not enough statements. got=%d",
          len(program.Statements))
  }
  stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
  if !ok {
      t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",
          program.Statements[0])
  }

  literal, ok := stmt.Expression.(*ast.IntegerLiteral)
  if !ok {
      t.Fatalf("exp not *ast.IntegerLiteral. got=%T", stmt.Expression)
  }
  if literal.Value != 5 {
      t.Errorf("literal.Value not %d. got=%d", 5, literal.Value)
  }
  if literal.TokenLiteral() != "5" {
      t.Errorf("literal.TokenLiteral not %s. got=%s", "5",
          literal.TokenLiteral())
  }
}
