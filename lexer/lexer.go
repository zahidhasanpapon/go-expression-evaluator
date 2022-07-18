package lexer

import (
	"fmt"
	"log"

	"github.com/zahidhasanpapon/go-expression-evaluator/token"
	"github.com/zahidhasanpapon/go-expression-evaluator/utils"
)

// Tokenize performs a lexical analysis.
func Tokenize(s string) []token.Token {
	runes := []rune(s)
	tokens := []token.Token{}

	for len(runes) > 0 {
		// Skip white space
		_, runes = utils.ReadWhile(runes, utils.IsWhiteSpace)

		if len(runes) == 0 {
			break
		}

		r, shifted := utils.ShiftRune(runes)
		runes = shifted

		if utils.IsNumber(r) {
			s, readed := utils.ReadWhile(runes, utils.IsNumber)

			runes = readed
			token := token.Token{
				Type:    token.Number,
				Literal: string(r) + s,
			}

			tokens = append(tokens, token)
		} else if utils.IsOperator(string(r)) {
			token := token.Token{
				Type:    token.OPE,
				Literal: string(r),
			}
			tokens = append(tokens, token)
		} else if string(r) == "." {
			currToken := tokens[len(tokens)-1]
			tokens = tokens[:len(tokens)-1]
			currToken.Literal += string(r)
			r, shifted := utils.ShiftRune(runes)
			runes = shifted
			if utils.IsNumber(r) {
				s, readed := utils.ReadWhile(runes, utils.IsNumber)

				runes = readed
				currToken.Literal = currToken.Literal + s + string(r)
				fmt.Println("curr token", currToken)
				tokens = append(tokens, currToken)
			}
			tokens = append(tokens, currToken)
		} else {
			log.Fatalf("Char %s not allowed", string(r))
		}
	}
	return tokens
}
