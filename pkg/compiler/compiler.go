package compiler

import (
	"fmt"
	"strconv"

	"github.com/AvoidMe/glua/pkg/ast"
	"github.com/AvoidMe/glua/pkg/lexer"
)

type Compiler struct {
}

func New() *Compiler {
	return &Compiler{}
}

func (self *Compiler) Compile(body []ast.Node) ([]Bytecode, error) {
	var result []Bytecode
	for _, v := range body {
		switch node := v.(type) {
		case ast.FunctionCall:
			result = append(result, Bytecode{
				Opcode: LOAD_NAME,
				Args:   []any{string(node.Name)},
			})
			args, err := self.Compile(node.Args)
			if err != nil {
				return nil, err
			}
			result = append(result, args...)
			result = append(result, Bytecode{
				Opcode: CALL,
				Args:   []any{1}, // TODO: calc call args somewhere
			})
		case ast.BinaryExpr:
			left, err := self.Compile([]ast.Node{node.Left})
			if err != nil {
				return nil, err
			}
			result = append(result, left...)
			right, err := self.Compile([]ast.Node{node.Right})
			if err != nil {
				return nil, err
			}
			result = append(result, right...)
			result = append(result, Bytecode{
				Opcode: BINARY_OP,
				Args:   []any{node.Operator},
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
			result = append(result, Bytecode{
				Opcode: LOAD_CONST,
				Args:   []any{val},
			})
		case ast.UnaryMinus:
			body, err := self.Compile([]ast.Node{node.Body})
			if err != nil {
				return nil, err
			}
			result = append(result, body...)
			result = append(result, Bytecode{
				Opcode: UNARY_MINUS,
			})
		default:
			return nil, fmt.Errorf("Unexpected node type: %T", node)
		}
	}
	return result, nil
}
