package lexer

import (
	"bufio"
	"errors"
	"io"
)

type Lexer struct {
	Tokens []Lexeme
}

func New() *Lexer {
	return &Lexer{}
}

func (self *Lexer) Consume(reader io.Reader, filename string) error {
	buf := bufio.NewReader(reader)

	for {
		b, err := buf.ReadByte()
		if err != nil {
			if errors.Is(err, io.EOF) {
				return nil
			}
			return err
		}
		switch {
		case b == '(':
			self.Tokens = append(self.Tokens, Lexeme{
				Type: OPEN_BRACKET,
				Val:  []byte{b},
			})
		case b == ')':
			self.Tokens = append(self.Tokens, Lexeme{
				Type: CLOSE_BRACKET,
				Val:  []byte{b},
			})
		case b == '+':
			self.Tokens = append(self.Tokens, Lexeme{
				Type: PLUS,
				Val:  []byte{b},
			})
		case b == '-':
			self.Tokens = append(self.Tokens, Lexeme{
				Type: MINUS,
				Val:  []byte{b},
			})
		case b == '*':
			self.Tokens = append(self.Tokens, Lexeme{
				Type: MULT,
				Val:  []byte{b},
			})
		case b == '/':
			self.Tokens = append(self.Tokens, Lexeme{
				Type: DIV,
				Val:  []byte{b},
			})
		case b == ' ':
			continue
		case b == '\n' || b == '\r':
			self.Tokens = append(self.Tokens, Lexeme{
				Type: NEW_LINE,
				Val:  []byte{b},
			})
		case b == '\'' || b == '"':
			val := []byte{b}
			for {
				bb, err := buf.ReadByte()
				if err != nil {
					return err
				}
				val = append(val, bb)
				if bb == val[0] {
					self.Tokens = append(self.Tokens, Lexeme{
						Type: STRING_LITERAL,
						Val:  val[1 : len(val)-1],
					})
					break
				}
			}
		case b >= '0' && b <= '9':
			val := []byte{b}
			for {
				bb, err := buf.ReadByte()
				if err != nil {
					if errors.Is(err, io.EOF) {
						self.Tokens = append(self.Tokens, Lexeme{
							Type: DIGIT_LITERAL,
							Val:  val,
						})
						return nil
					}
					return err
				}
				if bb >= '0' && bb <= '9' {
					val = append(val, bb)
				} else {
					self.Tokens = append(self.Tokens, Lexeme{
						Type: DIGIT_LITERAL,
						Val:  val,
					})
					err := buf.UnreadByte()
					if err != nil {
						return err
					}
					break
				}
			}
		default:
			val := []byte{b}
			for {
				bb, err := buf.ReadByte()
				if err != nil {
					if errors.Is(err, io.EOF) {
						self.Tokens = append(self.Tokens, Lexeme{
							Type: STRING,
							Val:  val,
						})
						return nil
					}
					return err
				}
				if !(bb == '(' ||
					bb == ')' ||
					bb == '+' ||
					bb == '-' ||
					bb == '*' ||
					bb == '/' ||
					bb == ' ' ||
					bb == '\n') {
					val = append(val, bb)
				} else {
					self.Tokens = append(self.Tokens, Lexeme{
						Type: STRING,
						Val:  val,
					})
					err := buf.UnreadByte()
					if err != nil {
						return err
					}
					break
				}
			}
		}
	}
}
