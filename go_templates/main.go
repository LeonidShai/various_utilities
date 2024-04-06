package main

import (
	"fmt"

	abstractfactory "go_templates/abstract_factory"
	adapter "go_templates/adapter"
	bridge "go_templates/bridge"
	builder "go_templates/builder"
	chainofresponsibility "go_templates/chain_of_responsibility"
	command "go_templates/command"
	composite "go_templates/composite"
	decorator "go_templates/decorator"
	facade "go_templates/facade"
	factorymethod "go_templates/factory_method"
	flyweight "go_templates/flyweight"
	interpreter "go_templates/interpreter"
	iterator "go_templates/iterator"
	mediator "go_templates/mediator"
	memento "go_templates/memento"
	observer "go_templates/observer"
	prototype "go_templates/prototype"
	proxy "go_templates/proxy"
	singleton "go_templates/singleton"
	state "go_templates/state"
	strategy "go_templates/strategy"
	templatemethod "go_templates/template_method"
	visitor "go_templates/visitor"
)

func main() {
	fmt.Println("go templates")

	// template Singleton
	singleton.SingletonWorker()

	// template Adapter
	adapter.AdapterWorker()

	// template Visitor
	visitor.VisitorWorker()

	// template Observer
	observer.ObserverWorker()

	// template Mediator
	mediator.MediatorWorker()

	// template Facade
	facade.FacadeWorker()

	// template Abstract Factory
	abstractfactory.AbstractFactoryWorker()

	// template Factory Method (Factory)
	factorymethod.FactoryMethodWorker()

	// template Builder
	builder.BuilderWorker()

	// template Prototype
	prototype.PrototypeWorker()

	// template Decorator
	decorator.DecoratorWorker()

	// template Bridge
	bridge.BridgeWorker()

	// template Composite
	composite.CompositeWorker()

	// template Memento
	memento.MementoWorker()

	// template Proxy
	proxy.ProxyWorker()

	// template Command
	command.CommandWorker()

	// template Template Method
	templatemethod.TemplateMethodWorker()

	// template State
	state.StateWorker()

	// template Chain of Responsibility
	chainofresponsibility.ChainOfResponsibilityWorker()

	// template Strategy
	strategy.StrategyWorker()

	// template Iterator
	iterator.IteratorWorker()

	// template Flyweight
	flyweight.FlyweightWorker()

	// template Interpreter
	interpreter.InterpreterWorker()
}
