package parser

import (
	"fmt"
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

func testIntegerLiteral(t *testing.T, expression ast.Expression, value int64) bool {
	integer, ok := expression.(*ast.IntegerLiteral)
	if !ok {
		t.Errorf("expression was not an ast.IntegerLiteral as expected. Got: %T", expression)
		return false
	}

	if integer.Value != value {
		t.Errorf("value of IntegerLiteral was not %d as expected. Got: %d", value, integer.Value)
		return false
	}

	if integer.TokenLiteral() != fmt.Sprintf("%d", value) {
		t.Errorf("integ.TokenLiteral not %d. got=%s", value, integer.TokenLiteral())
    return false
	}

	return true
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

func TestParsingPrefixExpressions(t *testing.T) {
	prefixTests := []struct{
		input string
		operator string
		integerValue int64
	}{
		{"!5", "!", 5},
		{"-15", "-", 15},
	}

	for _, testCase := range prefixTests {
		lexer := lexer.New(testCase.input)
		parser := New(lexer)
		program := parser.ParseProgram()
		checkParserErrors(t, parser)

		if len(program.Statements) != 1 {
			t.Fatalf("the programs statmentcount was not 1 as expected got: %d", len(program.Statements))
		}

		statement, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("the programs statement was not of type ExpressionStatement as expected/ Got: %T", statement.Expression)
		}

		expression, ok := statement.Expression.(*ast.PrefixExpression)
		if !ok {
			t.Fatalf("the expression was not of type PrefixExpression as expected/ Got: %T", statement.Expression)
		}

		if expression.Operator != testCase.operator {
			t.Fatalf("the expressions operator was not %s as expected. Got: %s", testCase.operator, expression.Operator)
		}

		if !testIntegerLiteral(t, expression.Right, testCase.integerValue) {
			return
		}
	}
}

func TestParsingInfixExpressions(t *testing.T) {
	infixTests := []struct{
		input string
		leftValue int64
		operator string
		rightValue int64
	}{
		{"5 + 5;", 5, "+", 5},
    {"5 - 5;", 5, "-", 5},
    {"5 * 5;", 5, "*", 5},
    {"5 / 5;", 5, "/", 5},
    {"5 > 5;", 5, ">", 5},
    {"5 < 5;", 5, "<", 5},
    {"5 == 5;", 5, "==", 5},
    {"5 != 5;", 5, "!=", 5},
	}

	for _, testCase := range infixTests {
		lexer := lexer.New(testCase.input)
		parser := New(lexer)
		program := parser.ParseProgram()
		checkParserErrors(t, parser)

		if len(program.Statements) != 1 {
			t.Fatalf("Expected only one statement got: %d", len(program.Statements))
		}

		statement, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("Expected first statement to be an ExpressionStatement. Got: %T", program.Statements[0])
		}

		expression, ok := statement.Expression.(*ast.InfixExpression)
		if !ok {
			t.Fatalf("Expected the expression to be an InfixExpression. Got: %T", statement.Expression)
		}

		if !testIntegerLiteral(t, expression.Left, testCase.leftValue) {
			return
		}

		if expression.Operator != testCase.operator {
			t.Fatalf("Expected operator %s. Got: %s", testCase.operator, expression.Operator)
		}


		if !testIntegerLiteral(t, expression.Right, testCase.rightValue) {
			return
		}
	}
}

func TestOperatorPrecedenceParsing(t *testing.T) {
	tests := []struct {
		input string
		expected string
	}{
	  {
	    "-a * b",
      "((-a) * b)",
    },
    {
      "!-a",
      "(!(-a))",
    },
    {
      "a + b + c",
      "((a + b) + c)",
    },
    {
      "a + b - c",
      "((a + b) - c)",
    },
    {
      "a * b * c",
      "((a * b) * c)",
    },
    {
      "a * b / c",
      "((a * b) / c)",
    },
    {
      "a + b / c",
      "(a + (b / c))",
    },
    {
      "a + b * c + d / e - f",
      "(((a + (b * c)) + (d / e)) - f)",
    },
    {
      "3 + 4; -5 * 5",
      "(3 + 4)((-5) * 5)",
    },
    {
      "5 > 4 == 3 < 4",
      "((5 > 4) == (3 < 4))",
    },
    {
      "5 < 4 != 3 > 4",
      "((5 < 4) != (3 > 4))",
    },
    {
      "3 + 4 * 5 == 3 * 1 + 4 * 5",
      "((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))",
    },
	}

	for _, testCase := range tests {
		lexer := lexer.New(testCase.input)
		parser := New(lexer)
		program := parser.ParseProgram()
		checkParserErrors(t, parser)

		actual := program.String()
		if actual != testCase.expected {
			t.Errorf("Expected=%q, got%q", testCase.expected, actual)
		}
	}
}
