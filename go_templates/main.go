package main

import (
	"fmt"
)

// type General struct {
// 	Chislo int
// }

// func (g *General) ChangeChislo(i int) {
// 	g.Chislo = i
// }

// type Concrete struct {
// 	General

// 	Name string
// }

func main() {
	fmt.Println("go templates")

	// template Singleton
	singletonWorker()

	// template Adapter
	adapterWorker()

	// template Visitor
	visitorWorker()

	// template Observer
	observerWorker()

	// template Mediator
	mediatorWorker()
}
