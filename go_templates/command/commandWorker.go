package command

import (
	"fmt"
)

func CommandWorker() {
	fmt.Println("---------- Command -----------")

	// создаем получателя (Receiver)
	lighter := NewDeviceTurnLight()
	// создаем объект Command и передаем наш девайс (Receiver)
	comm := &Presser{
		Device: lighter,
	}
	// создаем объект, который управляет включение/выключение (т.е. командой)
	invoker := &OnOffLight{
		Command: comm,
	}

	// выполняем команду
	invoker.TurnOnOffLight()
	invoker.TurnOnOffLight()

	fmt.Println("-------------------")
}
