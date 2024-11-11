package compiler

import (
	"fmt"
	"strconv"

	"github.com/AvoidMe/glua/pkg/ast"
	"github.com/AvoidMe/glua/pkg/lexer"
)

var (
	ENV_INDEX = 0
)

type Compiler struct {
	LocalRegister        int
	LocalRegisterMapping map[string]int
	Constants            []any
	ConstantMapping      map[any]int
}

func New() *Compiler {
	return &Compiler{
		LocalRegister:        0,
		LocalRegisterMapping: make(map[string]int),
		Constants:            make([]any, 0),
		ConstantMapping:      make(map[any]int),
	}
}

func (self *Compiler) Compile(body []ast.Node) (*Code, error) {
	var result []Bytecode
	for _, v := range body {
		switch node := v.(type) {
		case ast.FunctionCall:
			// Load function from upvalues
			const_mapping := self.AppendConst(string(node.Name))
			register_ref := self.LocalRegister
			self.LocalRegister++
			result = append(result, Bytecode{
				Opcode: OP_GETTABUP,
				Args:   [3]int{register_ref, ENV_INDEX, const_mapping},
			})

			// Compile arguments
			args, err := self.Compile(node.Args)
			if err != nil {
				return nil, err
			}
			result = append(result, args.Code...)

			// Call function
			result = append(result, Bytecode{
				Opcode: OP_CALL,
				// A -> func reference
				// B -> args count +1
				// C -> 1 (no return), 2 -> return?
				Args: [3]int{register_ref, 2, 1}, // TODO: calc call args somehow
			})
		case ast.BinaryExpr:
			left, err := self.Compile([]ast.Node{node.Left})
			if err != nil {
				return nil, err
			}
			result = append(result, left.Code...)
			right, err := self.Compile([]ast.Node{node.Right})
			if err != nil {
				return nil, err
			}
			result = append(result, right.Code...)
			var opcode Opcode
			switch node.Operator {
			case lexer.PLUS:
				opcode = OP_ADD
			case lexer.MINUS:
				opcode = OP_SUB
			case lexer.MULT:
				opcode = OP_MUL
			case lexer.DIV:
				opcode = OP_DIV
			default:
				return nil, err
			}
			self.LocalRegister--
			register_ref := self.LocalRegister - 1
			result = append(result, Bytecode{
				Opcode: opcode,
				Args:   [3]int{register_ref, register_ref, register_ref + 1},
			})
		case ast.Literal:
			var val any
			switch node.Type {
			case lexer.DIGIT_LITERAL:
				digit, err := strconv.Atoi(string(node.Val))
				if err != nil {
					return nil, err
				}
				val = digit
			case lexer.STRING_LITERAL:
				val = string(node.Val)
			}
			const_mapping := self.AppendConst(val)
			register_value := self.LocalRegister
			self.LocalRegister++
			result = append(result, Bytecode{
				Opcode: OP_LOADK,
				Args:   [3]int{register_value, const_mapping},
			})
		case ast.UnaryMinus:
			body, err := self.Compile([]ast.Node{node.Body})
			if err != nil {
				return nil, err
			}
			result = append(result, body.Code...)
			register_value := self.LocalRegister - 1
			result = append(result, Bytecode{
				Opcode: OP_UNM,
				Args:   [3]int{register_value, register_value},
			})
		default:
			return nil, fmt.Errorf("Unexpected node type: %T", node)
		}
	}
	return &Code{
		Code:      result,
		Constants: self.Constants,
	}, nil
}

func (self *Compiler) AppendConst(val any) int {
	const_mapping, ok := self.ConstantMapping[val]
	if ok {
		return const_mapping
	}
	const_mapping = len(self.Constants)
	self.Constants = append(self.Constants, val)
	self.ConstantMapping[val] = const_mapping
	return const_mapping
}
