package parser

import (
	"fmt"

	"github.com/AvoidMe/glua/pkg/ast"
	"github.com/AvoidMe/glua/pkg/lexer"
)

type Parser struct {
	Tokens []lexer.Lexeme
	Pos    int
}

func New(tokens []lexer.Lexeme) *Parser {
	return &Parser{
		Tokens: tokens,
		Pos:    0,
	}
}

func (self *Parser) Consume() ([]ast.Node, error) {
	var result []ast.Node
	for self.Peek() != nil {
		expr, err := self.ParseStmt()
		if err != nil {
			return nil, err
		}
		result = append(result, expr)
	}
	return result, nil
}

func (self *Parser) ParseStmt() (ast.Node, error) {
	var result []ast.Node
	for self.Peek() != nil {
		tok := self.NextToken()
		switch tok.Type {
		case lexer.NEW_LINE:
			if len(result) == 1 {
				return result[0], nil
			}
			continue
		// Right now only function call
		case lexer.STRING:
			next_tok := self.NextToken()
			for next_tok != nil && next_tok.Type == lexer.NEW_LINE {
				next_tok = self.NextToken()
			}
			if next_tok == nil {
				return nil, fmt.Errorf("Missing open bracket for function call: %v", *tok)
			}
			if next_tok.Type != lexer.OPEN_BRACKET {
				return nil, fmt.Errorf("Expecting open bracket, got: %s", *next_tok)
			}
			node := ast.FunctionCall{
				Name: tok.Val,
			}
			body, err := self.ParseSubexpr()
			if err != nil {
				return nil, err
			}
			node.Args = body
			result = append(result, node)
		case lexer.DIGIT_LITERAL, lexer.STRING_LITERAL:
			result = append(result, ast.Literal{
				Type: tok.Type,
				Val:  tok.Val,
			})
		case lexer.MINUS:
			if len(result) == 0 {
				right, err := self.ParseStmt()
				if err != nil {
					return nil, err
				}
				result = append(result, ast.UnaryMinus{
					Body: right,
				})
				continue
			}
			left := result[len(result)-1]
			result = result[:len(result)-1]
			right, err := self.ParseStmt()
			if err != nil {
				return nil, err
			}
			result = append(result, ast.BinaryExpr{
				Left:     left,
				Right:    right,
				Operator: tok.Type,
			})
		case lexer.PLUS, lexer.MULT, lexer.DIV:
			if len(result) == 0 {
				return nil, fmt.Errorf("Unexpected token: %v", *tok)
			}
			left := result[len(result)-1]
			result = result[:len(result)-1]
			right, err := self.ParseStmt()
			if err != nil {
				return nil, err
			}
			result = append(result, ast.BinaryExpr{
				Left:     left,
				Right:    right,
				Operator: tok.Type,
			})
		default:
			return nil, fmt.Errorf("Unexpected token: %s", *tok)
		}
	}
	if len(result) > 1 {
		return nil, fmt.Errorf("Internal error")
	}
	return result[0], nil
}

func (self *Parser) ParseSubexpr() ([]ast.Node, error) {
	arr := []lexer.Lexeme{}
	arr_tok := self.NextToken()
	for arr_tok != nil && arr_tok.Type != lexer.CLOSE_BRACKET {
		arr = append(arr, *arr_tok)
		arr_tok = self.NextToken()
	}
	if arr_tok == nil {
		return nil, fmt.Errorf("Unclosed bracket")
	}
	body, err := New(arr).Consume()
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (self *Parser) NextToken() *lexer.Lexeme {
	tok := self.Peek()
	if tok == nil {
		return tok
	}
	self.Pos++
	return tok
}

func (self *Parser) Peek() *lexer.Lexeme {
	if self.Pos == len(self.Tokens) {
		return nil
	}
	return &self.Tokens[self.Pos]
}
