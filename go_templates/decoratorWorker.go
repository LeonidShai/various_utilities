package main

import (
	"fmt"
	"go_templates/decorator"
)

func decoratorWorker() {
	fmt.Println("---------- Decorator -----------")

	simple := decorator.NewSimpleItem("chair")
	dec := decorator.NewDecorItem(&simple)
	fmt.Println(dec.GetName())
	dec.Print()

	// decoration function
	calcDouble := decorator.Wrapper(decorator.Double)
	_ = calcDouble(5)

	fmt.Println("-------------------")
}
