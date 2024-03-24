package chainofresponsibility

import "fmt"

// Шаблон для запуска выполнения последовательных действий (цепочки обязанностей)
// Допустим, есть несколько объектов(*), у которых нужно
// вызвать последовательно методы для выполнения каких-то действий
// (просто действий или действий над каким-то объектом)

// интерфейс для таких объектов(*)
type OrderHandler interface {
	SetNextHandler(OrderHandler)
	Handle(*Soup)
}

// структура, описывающая объект, над которым будут совершаться действия
type Soup struct {
	PourWaterPan     bool
	PeelVegetables   bool
	PutVegetablesPan bool
	PutSeasoning     bool
	Boil             bool
	Eat              bool
}

// структура, которая будет запускать всю цепочку (но можно было не заводить отдельную)
type CookSoupProcessor struct {
	handler OrderHandler
}

func (p *CookSoupProcessor) SetNextHandler(h OrderHandler) {
	p.handler = h
}

func (p *CookSoupProcessor) Handle(s *Soup) {
	fmt.Println("Start to cook a soup")
	p.handler.Handle(s)
}

// ниже структуры для каждого отдельного действия, удовлетворяющие интерфейсу
type PourWaterPanProcess struct {
	nextHandler OrderHandler
}

func (p *PourWaterPanProcess) SetNextHandler(h OrderHandler) {
	p.nextHandler = h
}

func (p *PourWaterPanProcess) Handle(s *Soup) {
	if p.nextHandler != nil {
		s.PourWaterPan = true
		fmt.Println("Pour water to pan")
		p.nextHandler.Handle(s)
	} else {
		fmt.Println("Uuups! A little problem!")
	}
}

type PeelVegetablesProcess struct {
	nextHandler OrderHandler
}

func (p *PeelVegetablesProcess) SetNextHandler(h OrderHandler) {
	p.nextHandler = h
}

func (p *PeelVegetablesProcess) Handle(s *Soup) {
	if p.nextHandler != nil {
		s.PeelVegetables = true
		fmt.Println("Peel vegetables for soup")
		p.nextHandler.Handle(s)
	} else {
		fmt.Println("Uuups! A little problem!")
	}
}

type PutVegetablesPanProcess struct {
	nextHandler OrderHandler
}

func (p *PutVegetablesPanProcess) SetNextHandler(h OrderHandler) {
	p.nextHandler = h
}

func (p *PutVegetablesPanProcess) Handle(s *Soup) {
	if p.nextHandler != nil {
		s.PutVegetablesPan = true
		fmt.Println("Put vegetables to pan")
		p.nextHandler.Handle(s)
	} else {
		fmt.Println("Uuups! A little problem!")
	}
}

type PutSeasoningProcess struct {
	nextHandler OrderHandler
}

func (p *PutSeasoningProcess) SetNextHandler(h OrderHandler) {
	p.nextHandler = h
}

func (p *PutSeasoningProcess) Handle(s *Soup) {
	if p.nextHandler != nil {
		s.PutSeasoning = true
		fmt.Println("Put seasoning to pan")
		p.nextHandler.Handle(s)
	} else {
		fmt.Println("Uuups! A little problem!")
	}
}

type BoilProcess struct {
	nextHandler OrderHandler
}

func (p *BoilProcess) SetNextHandler(h OrderHandler) {
	p.nextHandler = h
}

func (p *BoilProcess) Handle(s *Soup) {
	if p.nextHandler != nil {
		s.Boil = true
		fmt.Println("Boil the soup about 30 minutes")
		p.nextHandler.Handle(s)
	} else {
		fmt.Println("Uuups! A little problem!")
	}
}

type EatProcess struct {
	nextHandler OrderHandler
}

func (p *EatProcess) SetNextHandler(h OrderHandler) {
	p.nextHandler = h
}

func (p *EatProcess) Handle(s *Soup) {
	if s.Boil {
		s.Eat = true
		fmt.Println("Soup is ready! Eat the soup!")
	} else if p.nextHandler != nil {
		fmt.Println("Soup is not ready!")
		p.nextHandler.Handle(s)
	} else {
		fmt.Println("Uuups! A little problem!")
	}
}
