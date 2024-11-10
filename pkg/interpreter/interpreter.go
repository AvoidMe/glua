package interpreter

import (
	"fmt"

	"github.com/AvoidMe/glua/pkg/compiler"
	"github.com/AvoidMe/glua/pkg/lexer"
)

type Interpreter struct {
	Stack   []any
	Globals map[string]Function
}

func New() *Interpreter {
	return &Interpreter{
		Stack: make([]any, 0),
		Globals: map[string]Function{
			"print": PrintFunction,
		},
	}
}

func (self *Interpreter) Run(code []compiler.Bytecode) error {
	for _, v := range code {
		switch v.Opcode {
		case compiler.LOAD_CONST:
			self.Push(v.Args[0])
		case compiler.LOAD_NAME:
			self.Push(self.Globals[v.Args[0].(string)])
		case compiler.BINARY_OP:
			right := self.Pop().(int)
			left := self.Pop().(int)
			switch v.Args[0].(lexer.LexemeType) {
			case lexer.PLUS:
				self.Push(left + right)
			case lexer.MINUS:
				self.Push(left - right)
			case lexer.MULT:
				self.Push(left * right)
			case lexer.DIV:
				self.Push(left / right)
			}
		case compiler.UNARY_MINUS:
			val := self.Pop().(int)
			self.Push(-val)
		case compiler.CALL:
			argsCount := v.Args[0].(int)
			args := make([]any, argsCount)
			for i := 0; i < argsCount; i++ {
				args[argsCount-i-1] = self.Pop()
			}
			function := self.Pop().(Function)
			function(args)
		default:
			return fmt.Errorf("Unexpected opcode: %s", v.Opcode)
		}
	}
	return nil
}

func (self *Interpreter) Push(val any) {
	self.Stack = append(self.Stack, val)
}

func (self *Interpreter) Pop() any {
	if len(self.Stack) == 0 {
		panic("Pop on empty stack")
	}
	val := self.Stack[len(self.Stack)-1]
	self.Stack = self.Stack[:len(self.Stack)-1]
	return val
}
