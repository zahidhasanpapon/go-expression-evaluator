package lexer

import (
	"log"

	"github.com/zahidhasanpapon/go-expression-evaluator/token"
	"github.com/zahidhasanpapon/go-expression-evaluator/utils"
)

// Tokenize performs a lexical analysis.
func Tokenize(s string) []token.Token {
	runes := []rune(s)
	tokensQueue := []token.Token{}

	for len(runes) > 0 {
		// Skip white space
		_, runes = utils.ReadWhile(runes, utils.IsWhiteSpace)

		if len(runes) == 0 {
			break
		}

		r, shiftedRune := utils.ShiftRune(runes)
		runes = shiftedRune

		if utils.IsNumber(r) {
			s, readed := utils.ReadWhile(runes, utils.IsNumber)

			runes = readed
			token := token.Token{
				Type:    token.Number,
				Literal: string(r) + s,
			}

			tokensQueue = append(tokensQueue, token)
		} else if utils.IsOperator(string(r)) {

			token := token.Token{
				Type:    token.OPE,
				Literal: string(r),
			}

			tokensQueue = append(tokensQueue, token)
		} else if string(r) == "." {

			currToken := tokensQueue[len(tokensQueue)-1]
			tokensQueue = tokensQueue[:len(tokensQueue)-1]

			currToken.Literal += string(r)

			r, shiftedRune := utils.ShiftRune(runes)
			runes = shiftedRune

			if utils.IsNumber(r) {
				s, readed := utils.ReadWhile(runes, utils.IsNumber)

				runes = readed
				currToken.Literal = currToken.Literal + s + string(r)

				tokensQueue = append(tokensQueue, currToken)
			}
			tokensQueue = append(tokensQueue, currToken)
		} else {
			log.Fatalf("Char %s not allowed", string(r))
		}
	}
	return tokensQueue
}
