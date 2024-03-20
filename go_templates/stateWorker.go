package main

import (
	"fmt"
	"go_templates/state"
)

func stateWorker() {
	fmt.Println("---------- State -----------")
	machine := state.NewMachine()
	machine.StartMotion()
	machine.SwitchTransmission(1)
	machine.Gaz(7)
	machine.SwitchTransmission(2)
	machine.Gaz(11)
	machine.SwitchTransmission(2)
	machine.Gaz(13)
	machine.Velocity()
	machine.Brake()
	machine.Brake()
	machine.SwitchTransmission(1)
	machine.Brake()
	machine.StopMotion()

	fmt.Println("-------------------")
}
