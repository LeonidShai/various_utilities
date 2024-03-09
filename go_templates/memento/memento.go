package memento

// Паттерн для сохранения и загрузки состояния объекта
// Есть некая исходная структура (original) состояние которой может как-то изменяться во времени.
// Memento - структура для сохранения и восстановления состояний оригинальной структуры (original)
type memento struct {
	state string
}

// Caretaker (смотритель за снимками)
// позволяет сохранить список снимков и восстановить конкретный снимок по индексу
type caretaker struct {
	mementos []memento
}

func NewCaretaker() *caretaker {
	return &caretaker{
		mementos: make([]memento, 0),
	}
}

func (c *caretaker) AddMemento(m memento) {
	c.mementos = append(c.mementos, m)
}

func (c *caretaker) Memento(indx int) memento {
	return c.mementos[indx]
}
