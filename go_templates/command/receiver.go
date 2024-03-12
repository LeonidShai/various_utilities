package command

import "fmt"

// интерфейс ресивера
type Receiver interface {
	turnOnOff()
}

// конкретная структура, реализующая ресивер
type deviceTurnLight struct {
	isOn bool
}

func NewDeviceTurnLight() *deviceTurnLight {
	return &deviceTurnLight{
		isOn: false,
	}
}

func (d *deviceTurnLight) turnOnOff() {
	if d.isOn {
		d.turnOff()
	} else {
		d.turnOn()
	}
}

func (d *deviceTurnLight) turnOn() {
	d.isOn = true
	fmt.Println("turn on the light")
}

func (d *deviceTurnLight) turnOff() {
	d.isOn = false
	fmt.Println("turn off the light")
}
