package main

import (
	"fmt"
	"go_templates/flyweight"
)

func flyweightWorker() {
	fmt.Println("---------- Flyweight -----------")
	const (
		home   = "home"
		street = "street"

		black = "black"
		green = "green"
		blue  = "blue"
	)
	// создадим объекты (3 штуки платьев)
	homeBlack := flyweight.NewHomeDress(black)
	streetGreen := flyweight.NewStreetDress(green)
	streetBlue := flyweight.NewStreetDress(blue)

	// теперь создадим Фабрику наших платьев (хранилище)
	dressCloset := flyweight.NewDressClosetFactory(&homeBlack, &streetGreen, &streetBlue)

	// а теперь можем передавать нашу фабрику и использоватье её везде
	// тут в примере передавать не будем, будем просто использовать
	dress := dressCloset.GetDress(home, black)
	if dress != nil {
		dress.WhichDress()
	}

	dress = dressCloset.GetDress(home, blue)
	if dress != nil {
		dress.WhichDress()
	}

	dress = dressCloset.GetDress(street, blue)
	if dress != nil {
		dress.WhichDress()
	}

	dress = dressCloset.GetDress(street, green)
	if dress != nil {
		dress.WhichDress()
	}

	fmt.Println("-------------------")
}
