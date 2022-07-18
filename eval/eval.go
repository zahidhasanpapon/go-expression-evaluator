package eval

import (
	"log"
	"strconv"

	"github.com/zahidhasanpapon/go-expression-evaluator/token"
	"github.com/zahidhasanpapon/go-expression-evaluator/utils"
)

// Evaluate postfix expression
func Calculate(tokens []token.Token) float64 {
	result := 0.0

	stack := []float64{}

	for len(tokens) > 0 {
		t, shifted := utils.ShiftToken(tokens)
		tokens = shifted

		if t.Type == token.Number {
			value, err := strconv.ParseFloat(t.Literal, 64)
			if err != nil {
				log.Fatalf("cast error %s", t.Literal)
			}
			stack = append(stack, value)
		} else {
			// Case of OPE, get operands
			last, lastPoped := utils.Pop(stack)
			stack = lastPoped

			secondLast, secondLastPoped := utils.Pop(stack)
			stack = secondLastPoped

			switch t.Literal {
			case "+":
				calc := secondLast + last
				stack = append(stack, calc)
			case "-":
				calc := secondLast - last
				stack = append(stack, calc)
			case "*":
				calc := secondLast * last
				stack = append(stack, calc)
			case "/":
				calc := secondLast / last
				stack = append(stack, calc)
			}
		}
	}

	result, _ = utils.Pop(stack)

	return result
}
