package interpreter

import (
	"fmt"

	"github.com/AvoidMe/glua/pkg/compiler"
)

type Interpreter struct {
	Stack     []any
	Upvalues  []any
	Constants []any
}

func New() *Interpreter {
	_ENV := LuaTable{
		"print": LuaFunction(LuaPrintFunction),
	}
	return &Interpreter{
		Upvalues:  []any{_ENV},
		Stack:     make([]any, 0),
		Constants: make([]any, 0),
	}
}

func (self *Interpreter) Run(code *compiler.Code) error {
	self.Constants = code.Constants
	for _, v := range code.Code {
		switch v.Opcode {
		case compiler.OP_LOADK:
			self.Push(v.Args[0], self.Constants[v.Args[1]])
		case compiler.OP_GETTABUP:
			name := self.Constants[v.Args[2]]
			val := self.Upvalues[v.Args[1]].(LuaTable)
			self.Push(v.Args[0], val[name])
		case compiler.OP_ADD:
			left := self.Stack[v.Args[1]].(int)
			right := self.Stack[v.Args[2]].(int)
			self.Stack[v.Args[0]] = left + right
		case compiler.OP_SUB:
			left := self.Stack[v.Args[1]].(int)
			right := self.Stack[v.Args[2]].(int)
			self.Stack[v.Args[0]] = left - right
		case compiler.OP_MUL:
			left := self.Stack[v.Args[1]].(int)
			right := self.Stack[v.Args[2]].(int)
			self.Stack[v.Args[0]] = left * right
		case compiler.OP_DIV:
			left := self.Stack[v.Args[1]].(int)
			right := self.Stack[v.Args[2]].(int)
			self.Stack[v.Args[0]] = left / right
		case compiler.OP_UNM:
			val := self.Stack[v.Args[0]].(int)
			self.Stack[v.Args[1]] = -val
		case compiler.OP_CALL:
			val := self.Stack[v.Args[0]].(LuaFunction)
			args_begin := v.Args[0] + 1
			args_end := v.Args[0] + v.Args[1]
			args := self.Stack[args_begin:args_end]
			val(args)
		default:
			return fmt.Errorf("Unexpected opcode: %s", v.Opcode)
		}
	}
	return nil
}

func (self *Interpreter) Push(to int, val any) {
	if len(self.Stack) <= to {
		newStack := make([]any, to+1)
		copy(newStack, self.Stack)
		self.Stack = newStack
	}
	self.Stack[to] = val
}

func (self *Interpreter) Pop() any {
	if len(self.Stack) == 0 {
		panic("Pop on empty stack")
	}
	val := self.Stack[len(self.Stack)-1]
	self.Stack = self.Stack[:len(self.Stack)-1]
	return val
}
