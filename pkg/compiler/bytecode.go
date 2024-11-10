package compiler

import "fmt"

//go:generate stringer -type=Opcode
type Opcode int

const (
	CALL Opcode = iota
	UNARY_MINUS
	BINARY_OP
	LOAD_NAME
	LOAD_CONST
)

type Bytecode struct {
	Opcode Opcode
	Args   []any
}

func (self Bytecode) String() string {
	return fmt.Sprintf("%s: %v", self.Opcode, self.Args)
}
