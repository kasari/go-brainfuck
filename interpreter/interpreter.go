package interpreter

import (
	"fmt"

	"github.com/kasari/go-brainfuck/ast"
)

const MemorySize = 1000

type Interpreter struct {
	memory  [MemorySize]int
	pointer int
}

func New() *Interpreter {
	return &Interpreter{}
}

func (i *Interpreter) Exec(node ast.Node) {
	switch n := node.(type) {
	case *ast.Program:
		for _, stmt := range n.Statements {
			i.Exec(stmt)
		}
	case *ast.IncrementPointerStatement:
		i.pointer++
	case *ast.DecrementPointerStatement:
		i.pointer--
	case *ast.IncrementDataStatement:
		i.memory[i.pointer]++
	case *ast.DecrementDataStatement:
		i.memory[i.pointer]--
	case *ast.InputStatement:
		fmt.Scanf("%c", &i.memory[i.pointer])
	case *ast.OutputStatement:
		fmt.Printf("%c", i.memory[i.pointer])
	case *ast.WhileStatement:
		for i.memory[i.pointer] != 0 {
			for _, stmt := range n.Statements {
				i.Exec(stmt)
			}
		}
	}
}
