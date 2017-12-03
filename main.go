package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/kasari/go-brainfuck/interpreter"
	"github.com/kasari/go-brainfuck/lexer"
	"github.com/kasari/go-brainfuck/parser"
)

func main() {
	r := os.Stdin
	if len(os.Args) > 1 {
		f, err := os.Open(os.Args[1])
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		r = f
	}

	input, err := ioutil.ReadAll(r)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	l := lexer.New(string(input))
	p := parser.New(l)
	program := p.ParseProgram()

	i := interpreter.New()
	i.Exec(program)
}
