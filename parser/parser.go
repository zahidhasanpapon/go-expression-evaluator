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

// PostfixExpression return postfix expression
func PostfixExpression(infix []token.Token) []token.Token {
	operatorStack := []token.Token{}
	tokensQueue := []token.Token{}

	for len(infix) > 0 {
		t, shiftedRune := utils.ShiftToken(infix)
		infix = shiftedRune

		tokenPrecedence := PrecedenceOperators[t.Literal]

		// Get last operator from stack
		stackPrecedence := 0
		if len(operatorStack) > 0 {
			stackPrecedence = PrecedenceOperators[operatorStack[len(operatorStack)-1].Literal]
		}

		if t.Type == token.Number {
			tokensQueue = append(tokensQueue, t)
		} else if t.Type == token.OPE && len(operatorStack) == 0 {
			operatorStack = append(operatorStack, t)
		} else if t.Type == token.OPE && tokenPrecedence > stackPrecedence {
			operatorStack = append(operatorStack, t)
		} else if t.Type == token.OPE && t.Literal == "(" {
			operatorStack = append(operatorStack, t)
		} else if t.Type == token.OPE && t.Literal == ")" {
			var ope token.Token

			ope, poped := utils.PopToken(operatorStack)
			operatorStack = poped
			for ope.Literal != "(" && len(operatorStack) != 0 {
				tokensQueue = append(tokensQueue, ope)

				ope, poped = utils.PopToken(operatorStack)
				operatorStack = poped
			}
		} else {
			// Case of stack predence is bigger of token precedence
			for tokenPrecedence <= stackPrecedence {
				token, poped := utils.PopToken(operatorStack)
				operatorStack = poped

				tokensQueue = append(tokensQueue, token)

				stackPrecedence = 0
				if len(operatorStack) > 0 {
					stackPrecedence = PrecedenceOperators[operatorStack[len(operatorStack)-1].Literal]
				}
			}

			operatorStack = append(operatorStack, t)
		}
	}

	// Add remaing operators from stack to queue
	for len(operatorStack) > 0 {
		token, poped := utils.PopToken(operatorStack)
		operatorStack = poped

		tokensQueue = append(tokensQueue, token)
	}

	return tokensQueue
}
