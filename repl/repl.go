package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/skatsuta/monkey-interpreter/eval"
	"github.com/skatsuta/monkey-interpreter/lexer"
	"github.com/skatsuta/monkey-interpreter/object"
	"github.com/skatsuta/monkey-interpreter/parser"
)

const prompt = ">> "

// Start starts Monkey REPL.
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Print(prompt)
		if !scanner.Scan() {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := eval.Eval(program, env)
		if evaluated == nil {
			continue
		}

		io.WriteString(out, evaluated.Inspect())
		io.WriteString(out, "\n")
	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, msg)
		io.WriteString(out, "\n")
	}
}
