Expression Evaluation in Golang

Mathematical expression evaluator with the [Shunting Yard Algorithm] (https://brilliant.org/wiki/shunting-yard-algorithm/), implemented in Go.

## Available operators

| Operator | Operation | What it returns       |
| -------- | --------- | --------------------- |
| +        | x + y     | Sum of x and y        |
| -        | x - y     | Difference of x and y |
| \*       | x \* y    | Product of x and y    |
| /        | x / y     | Quotient of x and y   |
| (        |           | Open parenthesis      |
| )        |           | Closed parenthesis    |

### Flow

```text
Input -> Lexical Analysis -> Parsing (Shunting Yard) --> Evaluate --> Output
```

## Use

```shell
go run main.go
> ((1+2)*(4/2))
6
```
