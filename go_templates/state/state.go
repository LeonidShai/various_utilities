package state

import "fmt"

// Шаблон Состояние. Тут все просто, есть несколько состояний у объекта.
// в эти состояния объект может переходить в определенные моменты
// удобно, чтобы не писать кучу if-ов

// имеем интерфейс, который должно описывать каждое состояние
type State interface {
	StartMotion()
	Gaz(int)
	SwitchTransmission(int)
	Brake()
	StopMotion()
}

// Machine is ...
type machine struct {
	restState         State
	startState        State
	accelerationState State
	brakingState      State
	switcherState     State

	currentState State

	velocity     int
	transmission int
}

func NewMachine() *machine {
	machine := &machine{
		velocity:     0,
		transmission: 0,
	}

	restState := &restState{machine: machine}
	startState := &startState{machine: machine}
	accelerationState := &accelerationState{machine: machine}
	brakingState := &brakingState{machine: machine}
	switcherState := &switcherState{machine: machine}

	machine.restState = restState
	machine.startState = startState
	machine.accelerationState = accelerationState
	machine.brakingState = brakingState
	machine.switcherState = switcherState

	machine.currentState = restState

	return machine
}

func (m *machine) Velocity() {
	fmt.Printf("Current velocity: %d\n", m.velocity)
}

func (m *machine) setState(state State) {
	m.currentState = state
}

func (m *machine) StartMotion() {
	m.currentState.StartMotion()
}
func (m *machine) Gaz(v int) {
	m.currentState.Gaz(v)
}
func (m *machine) SwitchTransmission(t int) {
	m.currentState.SwitchTransmission(t)
}
func (m *machine) Brake() {
	m.currentState.Brake()
}
func (m *machine) StopMotion() {
	m.currentState.StopMotion()
}
