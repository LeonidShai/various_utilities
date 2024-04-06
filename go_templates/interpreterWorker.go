package main

import (
	"fmt"
	"go_templates/interpreter"
)

func interpreterWorker() {
	fmt.Println("---------- Iterpreter -----------")
	const (
		expr1 = "2 + 3 - 4 + 2"
		expr2 = "2 + 4"
		expr3 = "2 + 4 - 7"
	)

	// простой калькулятор
	parser := interpreter.NewParser()
	parser.Parse(expr1)
	res := parser.Result()
	fmt.Printf("%s = %d\n", expr1, res)

	parser.Parse(expr2)
	res = parser.Result()
	fmt.Printf("%s = %d\n", expr2, res)

	parser.Parse(expr3)
	res = parser.Result()
	fmt.Printf("%s = %d\n", expr3, res)

	fmt.Println("-------------------")
}
