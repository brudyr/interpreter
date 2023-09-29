package ast

import (
	"testing"

	"github.com/brudyr/go-interpreter/token"
)

func TestString(t *testing.T) {
  program := &Program{
    Statements: []Statement{
      &LetStatement{
        Token: token.Token{ Type: token.LET, Value: "let" },
        Name: &Identifier{ 
          Token: token.Token{ Type: token.IDENT, Value: "myVar" },
          Value: "myVar",
        },
        Value: &Identifier{
          Token: token.Token{ Type: token.IDENT, Value: "anotherVar" },
          Value: "anotherVar",
        },
      },
    },
  }

  expectedValue := "let myVar = anotherVar;"
  if program.String() != expectedValue {
    t.Errorf("programm.String() mismatch! expected: %q got: %q", expectedValue, program.String())
  }
}
