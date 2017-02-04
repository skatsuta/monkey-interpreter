package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/skatsuta/monkey-interpreter/eval"
	"github.com/skatsuta/monkey-interpreter/lexer"
	"github.com/skatsuta/monkey-interpreter/object"
	"github.com/skatsuta/monkey-interpreter/parser"
	"github.com/skatsuta/monkey-interpreter/repl"
)

func main() {
	// Start Monkey REPL
	if len(os.Args) == 1 {
		fmt.Println("This is the Monkey programming language!")
		fmt.Println("Feel free to type in commands")
		repl.Start(os.Stdin, os.Stdout)
		return
	}

	// Run a Monkey script
	runProgram(os.Args[1])
}

func runProgram(filename string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not read %s: %v\n", os.Args[1], err)
		os.Exit(1)
	}

	p := parser.New(lexer.New(string(data)))
	program := p.ParseProgram()
	if len(p.Errors()) > 0 {
		for _, e := range p.Errors() {
			fmt.Fprintln(os.Stderr, e)
		}
		os.Exit(1)
	}

	env := object.NewEnvironment()
	evaluated := eval.Eval(program, env)
	if evaluated == nil {
		return
	}

	io.WriteString(os.Stdout, evaluated.Inspect()+"\n")
}
