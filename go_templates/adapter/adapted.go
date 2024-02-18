package adapter

import "fmt"

// адаптируемые структуры с методами
type Bus struct{}

func (c *Bus) TravelByBus() {
	fmt.Println("Travel By Bus")
}

type Plane struct{}

func (p *Plane) TravelByPlane() {
	fmt.Println("Travel By Plane")
}

type Train struct{}

func (t *Train) TravelByTrain() {
	fmt.Println("Travel By Trane")
}
