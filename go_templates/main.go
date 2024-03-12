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

	// template Abstract Factory
	abstractFactoryWorker()

	// template Factory Method (Factory)
	factoryMethodWorker()

	// template Builder
	builderWorker()

	// template Prototype
	prototypeWorker()

	// template Decorator
	decoratorWorker()

	// template Bridge
	bridgeWorker()

	// template Composite
	compositeWorker()

	// template Memento
	mementoWorker()

	// template Proxy
	proxyWorker()

	// template Command
	commandWorker()
}
