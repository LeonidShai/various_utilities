package strategy

import "fmt"

// паттерн Стратегия. позволяет изменять алгоритмы,
// независимо от клиентов
// либо когда много разных реализаций,
// либо когда нужна последовательность из разных действий

// заведем интерфейс описывающий нашу стратегию и 2 класса удовлетворяющие ему
type OurStrategyToday interface {
	HardWork()
}

type ChairSitting struct{}

func (s *ChairSitting) HardWork() {
	fmt.Println("Very hard work to sit on chair")
}

type EatSoup struct{}

func (s *EatSoup) HardWork() {
	fmt.Println("Very hard work to eat a soup")
}

// структура, которая будет использовать разные стратегии
type WhatDo struct {
	OurStrategy OurStrategyToday
}

func (w *WhatDo) SetNewStrategy(s OurStrategyToday) {
	w.OurStrategy = s
}

func (w *WhatDo) ExecuteOurStrategy() {
	w.OurStrategy.HardWork()
}
