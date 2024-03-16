package main

import (
	"fmt"
	templatemethod "go_templates/template_method"
)

func templateMethodWorker() {
	fmt.Println("---------- Template Method -----------")
	// 2 разных варианта использования
	// 1 вариант: (создаем wayToWork с конкретным объектом внутри
	// и у него вызываем метод process)
	vasya := &templatemethod.Vasya{}
	fedya := &templatemethod.Fedya{}
	vasyaWayToWork := &templatemethod.WayToWork{
		Wtw: vasya,
	}
	fedyaWayToWork := &templatemethod.WayToWork{
		Wtw: fedya,
	}
	vasyaWayToWork.ProcessWayToWork()
	fmt.Println()
	fedyaWayToWork.ProcessWayToWork()
	fmt.Println()

	// 2 вариант: (создаем конкретный объект, у которых внутри определяем wayToWork,
	// внутри которого помещаем этот конкретный объект.
	// И у конкретного объекта вызываем process)
	petya := &templatemethod.Petya{}
	dasha := &templatemethod.Dasha{}
	petya.WayToWork = templatemethod.WayToWork{
		Wtw: petya,
	}
	dasha.WayToWork = templatemethod.WayToWork{
		Wtw: dasha,
	}
	petya.ProcessWayToWork()
	fmt.Println()
	dasha.ProcessWayToWork()

	fmt.Println("-------------------")
}
