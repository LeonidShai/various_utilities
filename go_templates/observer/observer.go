package observer

import "fmt"

type Observer interface {
	Update(string)
	GetName() string
}

type ConcreteObserver struct {
	Name string
}

func (o *ConcreteObserver) Update(state string) {
	fmt.Printf("New state %v - observed by %v\n", state, o.Name)
}

func (o *ConcreteObserver) GetName() string {
	return o.Name
}
