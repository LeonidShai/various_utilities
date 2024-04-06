package builder

import (
	"fmt"
)

func BuilderWorker() {
	fmt.Println("---------- Builder-----------")

	chairBuilder := NewChairBuilder()
	chair := chairBuilder.ChairLegs(3).ChairBack("blue").Sitting("metal").Build()
	fmt.Println(chair)

	offChB := NewOfficeChairBuilder()
	offChB.ChairBack("grey")
	director := NewDirector(offChB)
	offCh := director.BuildChair()
	fmt.Println(offCh)

	homeChB := NewHomeChairBuilder()
	homeChB.ChairLegs(0)
	director.SetBuilder(homeChB)
	fmt.Println(director.BuildChair())

	fmt.Println("-------------------")
}
