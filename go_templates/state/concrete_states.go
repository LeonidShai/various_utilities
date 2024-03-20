package state

import "fmt"

type restState struct {
	machine *machine
}

func (s *restState) StartMotion() {
	fmt.Println("Rest state: start Motion -- switch to start state")
	s.machine.setState(s.machine.startState)
}
func (s *restState) Gaz(_ int) {
	fmt.Println("Rest state: gaz")
}
func (s *restState) SwitchTransmission(_ int) {
	fmt.Println("Rest state: SwitchTransmission")
}
func (s *restState) Brake() {
	fmt.Println("Rest state: Brake")
}
func (s *restState) StopMotion() {
	fmt.Println("Rest state: StopMotion")
}

type startState struct {
	machine *machine
}

func (s *startState) StartMotion() {
	fmt.Println("startState: start Motion")
}
func (s *startState) Gaz(_ int) {
	fmt.Println("startState: gaz")
}
func (s *startState) SwitchTransmission(t int) {
	if t == 1 {
		s.machine.transmission = t
		fmt.Printf("startState: SwitchTransmission to %d -- switch to accelerationState\n", t)
		s.machine.setState(s.machine.accelerationState)
	} else {
		fmt.Printf("startState: SwitchTransmission to %d -- bad transmission\n", t)
	}
}
func (s *startState) Brake() {
	fmt.Println("startState: Brake")
}
func (s *startState) StopMotion() {
	fmt.Println("startState: StopMotion -- switched to rest state")
	s.machine.setState(s.machine.restState)
}

type accelerationState struct {
	machine *machine
}

func (s *accelerationState) StartMotion() {
	fmt.Println("accelerationState: StartMotion")
}
func (s *accelerationState) Gaz(v int) {
	s.machine.velocity = v
	if v >= 10 {
		fmt.Printf("accelerationState: Gaz %d -- switch to switcher state\n", v)
		s.machine.setState(s.machine.switcherState)

		return
	}
	fmt.Printf("accelerationState: Gaz %d\n", v)
}
func (s *accelerationState) SwitchTransmission(_ int) {
	fmt.Println("accelerationState: SwitchTransmission")
}
func (s *accelerationState) Brake() {
	fmt.Println("accelerationState: Brake -- switched to rest state")
	s.machine.setState(s.machine.restState)
}
func (s *accelerationState) StopMotion() {
	fmt.Println("accelerationState: StopMotion")
}

type brakingState struct {
	machine *machine
}

func (s *brakingState) StartMotion() {
	fmt.Println("brakingState: StartMotion")
}
func (s *brakingState) Gaz(_ int) {
	fmt.Println("brakingState: Gaz")
}
func (s *brakingState) SwitchTransmission(_ int) {
	fmt.Println("brakingState: SwitchTransmission")
}
func (s *brakingState) Brake() {
	s.machine.velocity -= 5
	s.machine.Velocity()
	if s.machine.velocity >= 7 {
		fmt.Println("brakingState: Brake -- switcher state")
		s.machine.setState(s.machine.switcherState)
	} else if s.machine.velocity <= 0 {
		fmt.Println("brakingState: Brake")
		s.machine.transmission = 0
	} else {
		fmt.Println("brakingState: Brake")
	}
}
func (s *brakingState) StopMotion() {
	fmt.Println("brakingState: StopMotion -- rest state")
	s.machine.setState(s.machine.restState)
}

type switcherState struct {
	machine *machine
}

func (s *switcherState) StartMotion() {
	fmt.Println("switcherState: StartMotion")
}
func (s *switcherState) Gaz(_ int) {
	fmt.Println("switcherState: Gaz")
}
func (s *switcherState) SwitchTransmission(t int) {
	if t == s.machine.transmission+1 {
		s.machine.transmission = t
		fmt.Printf("switcherState: SwitchTransmission %d -- switch to acceleration state\n", t)
		s.machine.setState(s.machine.accelerationState)
	} else if t == s.machine.transmission-1 {
		s.machine.transmission = t
		fmt.Printf("switcherState: SwitchTransmission %d -- switch to braking state\n", t)
		s.machine.setState(s.machine.brakingState)
	} else {
		fmt.Printf("switcherState: SwitchTransmission %d\n", t)
	}
}
func (s *switcherState) Brake() {
	fmt.Println("switcherState: Brake -- switch to braking state")
	s.machine.setState(s.machine.brakingState)
}
func (s *switcherState) StopMotion() {
	fmt.Println("switcherState: StopMotion")
}
