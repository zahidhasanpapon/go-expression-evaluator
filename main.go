package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/zahidhasanpapon/go-expression-evaluator/eval"
	"github.com/zahidhasanpapon/go-expression-evaluator/lexer"
	"github.com/zahidhasanpapon/go-expression-evaluator/parser"
)

func main() {
	for {
		fmt.Print("=> ")

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')

		infix := lexer.Tokenize(input)
		// fmt.Println(infix)
		tokens := parser.PostfixExpression(infix)
		// fmt.Println(tokens)
		reuslt := eval.Calculate(tokens)

		fmt.Println(reuslt)
	}
}
