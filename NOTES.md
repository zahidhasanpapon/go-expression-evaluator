Operators: +, -, \*, /
Numbers
Parenthesized expressions

1. Types of functions in go
2. make()
3. Infix
4. Operators
5. Operands

Shunting Yard Algorithm
We write mathematical expressions in infix notation. Operators have precedence and brackets override this precedence.
The shunting yard algorithm is a simple technique for parsing infix expressions containing binary operators of varying precedence. In general, the algorithm assigns to each operator its correct operands, taking into account the order of precedence. It can, therefore, be used to evaluate the expression immediately, to convert it into postfix, or to construct the corresponding syntax tree.

1. One stack for operations
2. One queue for the output
3. One array of tokens

Reverse Polish Notation

1. Expressions are parsed left to right
2. Each time a number or operand is read, we push it to the stack
3. Each time an operator comes up, we pop the required operands from the stack, perform the operations, and push the result back to the stack.
4. We are finished when there is no tokens to read. The final number on the stack is the result.

<!--  -->

1. First read shunting yard algorithm and understand
2. Undersand why this code is used
3. Implement data strucuteres locally
4. What is a lexer?
