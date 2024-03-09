package composite

import "fmt"

// Паттерн Компоновщик
// когда требуется иерархическая структура
// каждай узел - либо простой объект, либо группа объектов

type Composite interface {
	Search(string)
}

// Состовная структура со списком дочерних узлов, должна реализовывать интерфейс Компоновщика
type folder struct {
	components []Composite
	name       string
}

func NewFolder(name string) *folder {
	return &folder{
		components: make([]Composite, 0),
		name:       name,
	}
}

func (f *folder) Search(pattern string) {
	fmt.Printf("Search %s in folder %s\n", pattern, f.name)
	for _, component := range f.components {
		component.Search(pattern)
	}
}

func (f *folder) AddComponents(comps ...Composite) {
	f.components = append(f.components, comps...)
}

// Простая структура, которая также удовлетворяет интерфейсу Компоновщика
// и может быть в списке дочерних узлов составной структуры
type file struct {
	name string
}

func NewFile(name string) *file {
	return &file{
		name: name,
	}
}

func (f *file) Search(pattern string) {
	fmt.Printf("Search by %s in file %s\n", pattern, f.name)
}
