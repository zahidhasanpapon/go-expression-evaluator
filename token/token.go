// Idiomatic way to implement enumerated type
// https://yourbasic.org/golang/iota/

package token

type TypeInt int

const (
	Illegal TypeInt = iota
	EOF
	OpeningParenthesis
	ClosingParenthesis
	Plus
	Minus
	Divide
	Asterisk
	Number
)

type Token struct {
	Type    TypeInt
	Literal string
}

func NewToken(t TypeInt, l string) Token {
	return Token{Type: t, Literal: l}
}
