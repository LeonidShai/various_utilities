package main

import (
	"fmt"
	"go_templates/bridge"
)

func bridgeWorker() {
	fmt.Println("---------- Bridge -----------")
	// Работники Вася и Федя (абстракции конкретные)
	vasya := bridge.Vasya{}
	fedya := bridge.Fedya{}

	// Реализации работ (повар и водитель автобуса)
	cook := &bridge.Cook{}
	busDriver := &bridge.BusDriver{}

	// для каждой конкретной абстракции добавляем реализацию
	vasya.AddProfession(cook)
	fedya.AddProfession(busDriver)

	// выполнение
	vasya.Work()
	fedya.Work()

	fmt.Println("-------------------")
}
