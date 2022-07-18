package utils

import (
	"unicode"

	"github.com/zahidhasanpapon/go-expression-evaluator/token"
)

// ShiftToken shift token array
func ShiftToken(tokens []token.Token) (token.Token, []token.Token) {
	var t token.Token
	if len(tokens) > 0 {
		t = tokens[0]
		tokens = tokens[1:]
	}
	return t, tokens
}

// Pops from stack.
func Pop(values []float64) (float64, []float64) {
	var v float64
	if len(values) > 0 {
		v = values[len(values)-1]
		values = values[:len(values)-1]
	}
	return v, values
}

type isFunc func(rune) bool

// ReadWhile returns the first rune and the remaining rune
func ReadWhile(runes []rune, predicate isFunc) (string, []rune) {
	var s = ""
	for len(runes) > 0 && predicate(runes[0]) {
		c, shifted := ShiftRune(runes)
		runes = shifted
		s += string(c)
	}
	return s, runes
}

// ShftsRune
func ShiftRune(runes []rune) (rune, []rune) {
	var r rune
	if len(runes) > 0 {
		r = runes[0]
		runes = runes[1:]
	}
	return r, runes
}

// Checks if current rune is white space or not.
func IsWhiteSpace(r rune) bool {
	return unicode.IsSpace(r)
}

// Checks if r rune is number or not
func IsNumber(r rune) bool {
	return unicode.IsDigit(r)
}

func IsOperator(r string) bool {
	if r == "+" || r == "-" || r == "/" || r == "*" || r == "(" || r == ")" {
		return true
	}
	return false
}

func PopToken(tokens []token.Token) (token.Token, []token.Token) {
	var t token.Token
	if len(tokens) > 0 {
		t = tokens[len(tokens)-1]
		tokens = tokens[:len(tokens)-1]
	}
	return t, tokens
}
