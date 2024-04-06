package facade

import (
	"fmt"
)

type GetUpIf interface {
	GetUpPersonByLift() bool
}

func tryingToGetUpPersonByLift(fc GetUpIf) bool {
	return fc.GetUpPersonByLift()
}

func FacadeWorker() {
	fmt.Println("----------Facade -----------")

	fc := NewFacade()
	if tryingToGetUpPersonByLift(fc) {
		fmt.Printf("This is success. Lift can get up %v\n", fc.PersonName())
	} else {
		fmt.Printf("It is a bad case. Lift can't get up %v\n", fc.PersonName())
	}

	fmt.Println("-------------------")
}
