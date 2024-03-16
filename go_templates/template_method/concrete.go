package templatemethod

import "fmt"

type Vasya struct {
	WayToWork
}

func (v *Vasya) GetReadyToGo() {
	fmt.Println("Vasya: wake up, eat breakfast")
}

func (v *Vasya) GoByTransport() {
	fmt.Println("Vasya: go by car")
}

func (v *Vasya) FindFreeChair() {
	fmt.Println("Vasya: chair near the window")
}

type Fedya struct {
	WayToWork
}

func (f *Fedya) GetReadyToGo() {
	fmt.Println("Fedya: wake up")
}

func (f *Fedya) GoByTransport() {
	fmt.Println("Fedya: go by subway")
}

func (f *Fedya) FindFreeChair() {
	fmt.Println("Fedya: chair near the coffemachine")
}

type Petya struct {
	WayToWork
}

func (p *Petya) GetReadyToGo() {
	fmt.Println("Petya: wake up, drink coffee")
}

func (p *Petya) GoByTransport() {
	fmt.Println("Petya: go by bus")
}

func (p *Petya) FindFreeChair() {
	fmt.Println("Petya: chair near the Vasya")
}

type Dasha struct {
	WayToWork
}

func (d *Dasha) GetReadyToGo() {
	fmt.Println("Dasha: wake up, eat breakfast, drink coffee")
}

func (d *Dasha) GoByTransport() {
	fmt.Println("Dasha: go by bus and subway")
}

func (d *Dasha) FindFreeChair() {
	fmt.Println("Dasha: chair near the Petya")
}
