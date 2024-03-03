package bridge

import "fmt"

// реализация: конкретные профессии
type Worker interface {
	DoWork()
}

type Cook struct{}

func (c *Cook) DoWork() {
	fmt.Println("cook food")
}

type BusDriver struct{}

func (b *BusDriver) DoWork() {
	fmt.Println("drive a bus")
}
