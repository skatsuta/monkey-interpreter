package main

import (
	"fmt"
	"os"

	"github.com/skatsuta/monkey-interpreter/repl"
)

func main() {
	fmt.Printf("This is the Monkey programming language!\n")
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
