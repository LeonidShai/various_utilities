package facade

// суть Фасада - объединить и скрыть внутреннюю реализацию, состоящую из множества структур и действий над ними
// предоставив интерфейс пользователю (одну или несколько общих функций) для вызова выполнения всех операций.
type Facade struct {
	Person *Person
	Lift   *Lift
}

func NewFacade() *Facade {
	return &Facade{
		Person: NewPerson("Vasya", 64),
		Lift:   NewLift(50),
	}
}

func (f *Facade) GetUpPersonByLift() bool {
	if f.Lift.GetUp(f.Person) {
		return true
	}
	f.Person.ScreamAaaa()

	return false
}

func (f *Facade) PersonName() (name string) {
	return f.Person.Name
}
