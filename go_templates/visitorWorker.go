package main

import (
	"fmt"

	visitor "go_templates/visitor"
)

func visitorWorker() {
	fmt.Println("---------- Visitor -----------")

	person := visitor.NewPerson("Vasya")
	drinkerTea := visitor.TeaDrinker{}
	person.Accept(&drinkerTea)

	eatingFood := visitor.FoodEating{}
	person.Accept(&eatingFood)

	sleepDream := visitor.DreamSeeing{}
	person.Accept(&sleepDream)

	fmt.Println("-------------------")
}
