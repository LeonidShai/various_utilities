package templatemethod

import "fmt"

// определяет основу алгоритма, позволяет подклассам переопределить
// некоторые шаги алгоритма, не изменяя структуры

// интерфейс, которому должны будут удовлетворять классы,
// которые хотим подстроить под один алгоритм
type WayToWorkIf interface {
	GetReadyToGo()
	GoByTransport()
	FindFreeChair()
}

// некоторая структура у которой есть метод с алгоритмом действий
type WayToWork struct {
	Wtw WayToWorkIf
}

func (wtw *WayToWork) Voila() {
	fmt.Println("WayToWork: Voila! You are at Work")
}

func (wtw *WayToWork) ProcessWayToWork() {
	wtw.Wtw.GetReadyToGo()
	wtw.Wtw.GoByTransport()
	wtw.Wtw.FindFreeChair()
	wtw.Voila()
}
