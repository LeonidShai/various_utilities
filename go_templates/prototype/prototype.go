package prototype

type item struct {
	name string
}

func NewItem(name string) item {
	return item{
		name: name,
	}
}

// Prototype interface паттерн для клонирования/копирования объекта
// для создания точно такого же, но другого
// (вот точно такой же миной, только больше и другой)
type ClonerIf interface {
	GetName() string

	Clone() ClonerIf
}

func (i *item) GetName() string {
	return i.name
}

func (i *item) Clone() ClonerIf {
	return &item{
		name: i.name,
	}
}
