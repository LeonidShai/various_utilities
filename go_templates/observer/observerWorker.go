package observer

import (
	"fmt"
)

func ObserverWorker() {
	fmt.Println("---------- Observer -----------")
	stater := NewStater()

	vasya := &ConcreteObserver{
		Name: "Vasya",
	}
	petya := &ConcreteObserver{
		Name: "Petya",
	}

	stater.RegisterObserver(vasya)
	stater.RegisterObserver(petya)

	stater.ChangeState("First state")

	stater.DeregisterObserver(vasya)
	stater.ChangeState("Second state")

	fmt.Println("-------------------")
}
