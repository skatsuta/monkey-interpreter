package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/skatsuta/monkey-interpreter/lexer"
	"github.com/skatsuta/monkey-interpreter/token"
)

const prompt = ">> "

// Start starts Monkey REPL.
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(prompt)
		if !scanner.Scan() {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
