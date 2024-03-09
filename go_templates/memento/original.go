package memento

import "fmt"

// Для сохранения и восстановления заведем MementoIf
type MementoIf interface {
	Save() memento
	Restore(memento)
}

type original struct {
	state string
}

func NewOriginal() *original {
	return &original{
		state: "initial",
	}
}

func (o *original) Save() memento {
	fmt.Printf("Save %s state\n", o.state)

	return memento{
		state: o.state,
	}
}

func (o *original) Restore(m memento) {
	fmt.Printf("Restore %s state\n", m.state)
	o.state = m.state
}

func (o *original) ChangeState(state string) {
	fmt.Printf("Change %s state to %s state\n", o.state, state)
	o.state = state
}

func (o *original) PrintState() {
	fmt.Printf("Current state is %s\n", o.state)
}
