package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/brudyr/go-interpreter/lexer"
	"github.com/brudyr/go-interpreter/token"
)

const PROMT = ">>"

func Start(in io.Reader, out io.Writer) {
  scanner := bufio.NewScanner(in)

  for {
    fmt.Fprintf(out, PROMT)
    scanned := scanner.Scan()

    if !scanned {
      return
    }

    line := scanner.Text()
    l := lexer.New(line) 

    for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
      fmt.Fprintf(out, "%+v\n", tok)
    }
  }
}
