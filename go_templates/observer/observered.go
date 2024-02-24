package observer

import "fmt"

type Notifier interface {
	RegisterObserver(o Observer)
	DeregisterObserver(o Observer)
	Notify()
}

type stater struct {
	observers map[Observer]struct{}
	state     string
}

func NewStater() *stater {
	return &stater{
		observers: make(map[Observer]struct{}),
		state:     "Initial state",
	}
}

func (s *stater) RegisterObserver(observer Observer) {
	fmt.Printf("Registered %v\n", observer.GetName())
	s.observers[observer] = struct{}{}
}

func (s *stater) DeregisterObserver(observer Observer) {
	fmt.Printf("Deregistered %v\n", observer.GetName())
	delete(s.observers, observer)
}

func (s *stater) notify() {
	for observer := range s.observers {
		observer.Update(s.state)
	}
}

func (s *stater) ChangeState(newState string) {
	fmt.Printf("Change state from %v to %v\n", s.state, newState)
	s.state = newState
	s.notify()
}
