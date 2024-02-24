package main

import (
	"fmt"
	"go_templates/observer"
)

func observerWorker() {
	fmt.Println("---------- Observer -----------")
	stater := observer.NewStater()

	vasya := &observer.ConcreteObserver{
		Name: "Vasya",
	}
	petya := &observer.ConcreteObserver{
		Name: "Petya",
	}

	stater.RegisterObserver(vasya)
	stater.RegisterObserver(petya)

	stater.ChangeState("First state")

	stater.DeregisterObserver(vasya)
	stater.ChangeState("Second state")

	fmt.Println("-------------------")
}
