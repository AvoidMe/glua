package ast

import (
	"fmt"
	"strings"

	"github.com/AvoidMe/glua/pkg/lexer"
)

type Node interface {
	String() string
}

type FunctionCall struct {
	Name []byte
	Args []Node
}

func (self FunctionCall) String() string {
	args := []string{}
	for _, v := range self.Args {
		args = append(args, v.String())
	}
	return fmt.Sprintf("%s(%s)", string(self.Name), strings.Join(args, ", "))
}

type Literal struct {
	Type lexer.LexemeType
	Val  []byte
}

func (self Literal) String() string {
	return fmt.Sprintf("%s", self.Val)
}

type BinaryExpr struct {
	Left     Node
	Right    Node
	Operator lexer.LexemeType
}

func (self BinaryExpr) String() string {
	var lexeme string
	switch self.Operator {
	case lexer.PLUS:
		lexeme = "+"
	case lexer.MINUS:
		lexeme = "-"
	case lexer.MULT:
		lexeme = "*"
	case lexer.DIV:
		lexeme = "/"
	}
	return fmt.Sprintf("%s %s %s", self.Left, lexeme, self.Right)
}

type UnaryMinus struct {
	Body Node
}

func (self UnaryMinus) String() string {
	return fmt.Sprintf("-%s", self.Body)
}
