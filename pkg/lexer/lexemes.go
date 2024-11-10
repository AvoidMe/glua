package lexer

import "fmt"

//go:generate stringer -type=LexemeType
type LexemeType int

const (
	STRING LexemeType = iota
	DIGIT_LITERAL
	STRING_LITERAL
	OPEN_BRACKET
	CLOSE_BRACKET
	PLUS
	MINUS
	MULT
	DIV
	NEW_LINE
)

type Lexeme struct {
	Type LexemeType
	Val  []byte
}

func (self Lexeme) String() string {
	return fmt.Sprintf("%s: %s", self.Type.String(), string(self.Val))
}
