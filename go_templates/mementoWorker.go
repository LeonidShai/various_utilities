package main

import (
	"fmt"
	"go_templates/memento"
)

func mementoWorker() {
	fmt.Println("---------- Memento -----------")
	// Создадим оригинальный объект
	orig := memento.NewOriginal()
	origCaretaker := memento.NewCaretaker()
	// Сохраним и будем изменять и сохранять оригинальный объект
	origCaretaker.AddMemento(orig.Save())

	orig.ChangeState("first")
	origCaretaker.AddMemento(orig.Save())

	orig.ChangeState("second")
	origCaretaker.AddMemento(orig.Save())

	orig.PrintState()
	// Восстановим разные состояния исходного объекта
	orig.Restore(origCaretaker.Memento(1))
	orig.PrintState()

	orig.Restore(origCaretaker.Memento(0))
	orig.PrintState()

	fmt.Println("-------------------")
}
