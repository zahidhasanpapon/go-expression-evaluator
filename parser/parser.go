package parser

import (
	"github.com/zahidhasanpapon/go-expression-evaluator/token"
	"github.com/zahidhasanpapon/go-expression-evaluator/utils"
)

// PrecedenceOperators define Precedence of operators
var PrecedenceOperators = map[string]int{
	"(": 10,
	"+": 20,
	"-": 20,
	"/": 30,
	"*": 30,
}

// PostFixExpression return postfix expression
func PostfixExpression(infix []token.Token) []token.Token {
	opeStack := []token.Token{}
	tokens := []token.Token{}

	for len(infix) > 0 {
		t, shifted := utils.ShiftToken(infix)
		infix = shifted

		tokenPrecedence := PrecedenceOperators[t.Literal]

		// Get last OPE from stack
		stackPrecedence := 0
		if len(opeStack) > 0 {
			stackPrecedence = PrecedenceOperators[opeStack[len(opeStack)-1].Literal]
		}

		if t.Type == token.Number {
			// Case of NUM
			tokens = append(tokens, t)
		} else if t.Type == token.OPE && len(opeStack) == 0 {
			// Case First OPE in stack
			opeStack = append(opeStack, t)
		} else if t.Type == token.OPE && tokenPrecedence > stackPrecedence {
			// Case of token predecende is bigger of stack precedence
			opeStack = append(opeStack, t)
		} else if t.Type == token.OPE && t.Literal == "(" {
			// Case of open bracket
			opeStack = append(opeStack, t)
		} else if t.Type == token.OPE && t.Literal == ")" {
			// Case of closed bracket
			var ope token.Token

			ope, poped := utils.PopToken(opeStack)
			opeStack = poped
			for ope.Literal != "(" && len(opeStack) != 0 {
				tokens = append(tokens, ope)

				ope, poped = utils.PopToken(opeStack)
				opeStack = poped
			}
		} else {
			// Case of stack predence is bigger of token precedence
			for tokenPrecedence <= stackPrecedence {
				token, poped := utils.PopToken(opeStack)
				opeStack = poped

				tokens = append(tokens, token)

				stackPrecedence = 0
				if len(opeStack) > 0 {
					stackPrecedence = PrecedenceOperators[opeStack[len(opeStack)-1].Literal]
				}
			}

			opeStack = append(opeStack, t)
		}
	}

	// Add ope stack
	for len(opeStack) > 0 {
		token, poped := utils.PopToken(opeStack)
		opeStack = poped

		tokens = append(tokens, token)
	}

	return tokens
}
