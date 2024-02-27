package main

import (
	"fmt"
)

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

	// template Facade
	facadeWorker()
}
