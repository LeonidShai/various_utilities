package iterator

// Паттерн итератор - это внезапно итератор
// способ последовательного обращения ко всем элементам составного объекта
// без раскрытия внутреннего представления

// структура для объектов, которые будем итерировать
type Book struct {
	Name   string
	Author string
}

// интерфейс для итератора
type IteratorIf interface {
	HasNext() bool
	Next() *Book
}

// сам итератор
type Iterator struct {
	index int
	books []*Book
}

func (it *Iterator) HasNext() bool {
	if it.index < len(it.books) {
		return true
	}

	return false
}

func (it *Iterator) Next() *Book {
	if it.HasNext() {
		book := it.books[it.index]
		it.index++

		return book
	}

	return nil
}

// интерфейс для создания итераторов
type CollectionIf interface {
	CreateIterator() IteratorIf
}

type Collection struct {
	Books []*Book
}

func (c *Collection) CreateIterator() IteratorIf {
	return &Iterator{
		books: c.Books,
	}
}
