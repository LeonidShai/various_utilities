package templatemethod

import (
	"fmt"
)

func TemplateMethodWorker() {
	fmt.Println("---------- Template Method -----------")
	// 2 разных варианта использования
	// 1 вариант: (создаем wayToWork с конкретным объектом внутри
	// и у него вызываем метод process)
	vasya := &Vasya{}
	fedya := &Fedya{}
	vasyaWayToWork := &WayToWork{
		Wtw: vasya,
	}
	fedyaWayToWork := &WayToWork{
		Wtw: fedya,
	}
	vasyaWayToWork.ProcessWayToWork()
	fmt.Println()
	fedyaWayToWork.ProcessWayToWork()
	fmt.Println()

	// 2 вариант: (создаем конкретный объект, у которых внутри определяем wayToWork,
	// внутри которого помещаем этот конкретный объект.
	// И у конкретного объекта вызываем process)
	petya := &Petya{}
	dasha := &Dasha{}
	petya.WayToWork = WayToWork{
		Wtw: petya,
	}
	dasha.WayToWork = WayToWork{
		Wtw: dasha,
	}
	petya.ProcessWayToWork()
	fmt.Println()
	dasha.ProcessWayToWork()

	fmt.Println("-------------------")
}
