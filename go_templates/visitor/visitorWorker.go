package visitor

import (
	"fmt"
)

func VisitorWorker() {
	fmt.Println("---------- Visitor -----------")

	person := NewPerson("Vasya")
	drinkerTea := TeaDrinker{}
	person.Accept(&drinkerTea)

	eatingFood := FoodEating{}
	person.Accept(&eatingFood)

	sleepDream := DreamSeeing{}
	person.Accept(&sleepDream)

	fmt.Println("-------------------")
}
