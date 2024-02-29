package main

import (
	"fmt"
	"go_templates/builder"
)

func builderWorker() {
	fmt.Println("---------- Builder-----------")

	chairBuilder := builder.NewChairBuilder()
	chair := chairBuilder.ChairLegs(3).ChairBack("blue").Sitting("metal").Build()
	fmt.Println(chair)

	offChB := builder.NewOfficeChairBuilder()
	offChB.ChairBack("grey")
	director := builder.NewDirector(offChB)
	offCh := director.BuildChair()
	fmt.Println(offCh)

	homeChB := builder.NewHomeChairBuilder()
	homeChB.ChairLegs(0)
	director.SetBuilder(homeChB)
	fmt.Println(director.BuildChair())

	fmt.Println("-------------------")
}
