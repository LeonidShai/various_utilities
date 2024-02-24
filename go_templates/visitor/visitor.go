// Паттерн, который позволяет добавлять новые операции над объектом,
// без изменения самого объекта
package visitor

import "fmt"

// интерфейс для добавления разных операций к объекту
type visitorIf interface {
	VisitPerson(p *person)
}

// структуры с методами, которые требуется добавить к объекту
type TeaDrinker struct{}

func (td *TeaDrinker) VisitPerson(p *person) {
	fmt.Printf("Time to drink tea! %v pours a cup of tea.\n", p.name)
}

type FoodEating struct{}

func (fe *FoodEating) VisitPerson(p *person) {
	fmt.Printf("Time to eat some food! %v decided to fill his stomach.\n", p.name)
}

type DreamSeeing struct{}

func (ds *DreamSeeing) VisitPerson(p *person) {
	fmt.Printf("Time to sleep! %v is tired and goes to bad.\n", p.name)
}
