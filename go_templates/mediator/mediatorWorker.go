package mediator

import (
	"fmt"
)

func MediatorWorker() {
	fmt.Println("---------- Mediator -----------")
	md := &ConcreteMediator{}
	participant1 := Participant1{
		Participant: Participant{
			Mediator: md,
			Name:     "Participant1",
		},
		Id: 0,
	}
	participant2 := Participant2{
		Participant: Participant{
			Mediator: md,
			Name:     "Participant2",
		},
		Data: "some data",
	}
	md.RegisterParticipants(&participant1, &participant2)
	participant1.ChangeID(3)
	participant2.ChangeData("new data")

	fmt.Println("-------------------")
}
