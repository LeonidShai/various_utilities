package decorator

import (
	"fmt"
)

func DecoratorWorker() {
	fmt.Println("---------- Decorator -----------")

	simple := NewSimpleItem("chair")
	dec := NewDecorItem(&simple)
	fmt.Println(dec.GetName())
	dec.Print()

	// decoration function
	calcDouble := Wrapper(Double)
	_ = calcDouble(5)

	fmt.Println("-------------------")
}
