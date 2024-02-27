package facade

import "fmt"

type Lift struct {
	Weight int
}

func NewLift(weight int) *Lift {
	return &Lift{
		Weight: weight,
	}
}

func (l *Lift) GetUp(person *Person) bool {
	return l.Weight > person.Weight
}

type Person struct {
	Name   string
	Weight int
}

func NewPerson(name string, weight int) *Person {
	var temp = struct {
		tempName   string
		tempWeight int
	}{
		tempName:   name,
		tempWeight: weight,
	}

	return &Person{
		Name:   temp.tempName,
		Weight: temp.tempWeight,
	}
}

func (p *Person) ScreamAaaa() {
	fmt.Println("Aaaaa, I don't want go upstairs")
}
