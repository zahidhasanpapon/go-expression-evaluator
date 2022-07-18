package token

// TypeInt a lexer type.
type TypeInt int

const (
	OPE TypeInt = iota
	// Illegal TypeInt = iota
	// EOF
	// OpeningParenthesis
	// ClosingParenthesis
	// Plus
	// Minus
	// Divide
	// Asterisk
	Number
)

// Token define a token.
type Token struct {
	Type    TypeInt
	Literal string
}
