package visitor

// объект, к которому нужно добавить новую операцию
type person struct {
	name string
}

func NewPerson(name string) *person {
	return &person{
		name: name,
	}
}

// метод для добавления/выполнения новой операции
func (p *person) Accept(v visitorIf) {
	v.VisitPerson(p)
}
