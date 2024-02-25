package main

import (
	"fmt"
	mediator "go_templates/mediator"
)

func mediatorWorker() {
	fmt.Println("---------- Mediator -----------")
	md := &mediator.ConcreteMediator{}
	participant1 := mediator.Participant1{
		Mediator: md,
		Id:       0,
		Name:     "Participant1",
	}
	participant2 := mediator.Participant2{
		Mediator: md,
		Data:     "some data",
		Name:     "Participant2",
	}
	md.RegisterParticipants(&participant1, &participant2)
	participant1.ChangeID(2)
	participant2.ChangeData("new data")

	fmt.Println("-------------------")
}
