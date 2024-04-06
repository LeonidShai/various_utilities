package bridge

import (
	"fmt"
)

func BridgeWorker() {
	fmt.Println("---------- Bridge -----------")
	// Работники Вася и Федя (абстракции конкретные)
	vasya := Vasya{}
	fedya := Fedya{}

	// Реализации работ (повар и водитель автобуса)
	cook := &Cook{}
	busDriver := &BusDriver{}

	// для каждой конкретной абстракции добавляем реализацию
	vasya.AddProfession(cook)
	fedya.AddProfession(busDriver)

	// выполнение
	vasya.Work()
	fedya.Work()

	fmt.Println("-------------------")
}
