package main

import (
	"fmt"
	"go_templates/command"
)

func commandWorker() {
	fmt.Println("---------- Command -----------")

	// создаем получателя (Receiver)
	lighter := command.NewDeviceTurnLight()
	// создаем объект Command и передаем наш девайс (Receiver)
	comm := &command.Presser{
		Device: lighter,
	}
	// создаем объект, который управляет включение/выключение (т.е. командой)
	invoker := &command.OnOffLight{
		Command: comm,
	}

	// выполняем команду
	invoker.TurnOnOffLight()
	invoker.TurnOnOffLight()

	fmt.Println("-------------------")
}
