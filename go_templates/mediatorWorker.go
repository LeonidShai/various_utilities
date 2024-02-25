package main

import (
	"fmt"
	mediator "go_templates/mediator"
)

func mediatorWorker() {
	fmt.Println("---------- Mediator -----------")
	md := &mediator.ConcreteMediator{}
	participant1 := mediator.Participant1{
		Participant: mediator.Participant{
			Mediator: md,
			Name:     "Participant1",
		},
		Id: 0,
	}
	participant2 := mediator.Participant2{
		Participant: mediator.Participant{
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
