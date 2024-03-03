package bridge

import "fmt"

// Разделение абстракции и реализации
// Абстракция содержит ссылку на реализацию
// Абстракция (например, компьютер, который может вывести на печать документ)
// Реализация - принтер

// абстракция: человек с профессией
type Human interface {
	AddProfession(Worker)
	Work()
}

// конкретные абстракции
type Vasya struct {
	worker Worker
}

func (v *Vasya) AddProfession(w Worker) {
	v.worker = w
}

func (v *Vasya) Work() {
	fmt.Println("Vasya does his work")
	v.worker.DoWork()
}

type Fedya struct {
	worker Worker
}

func (v *Fedya) AddProfession(w Worker) {
	v.worker = w
}

func (v *Fedya) Work() {
	fmt.Println("Fedya does his work")
	v.worker.DoWork()
}
