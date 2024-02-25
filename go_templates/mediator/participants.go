package mediator

import "fmt"

type ParticipantIf interface {
	notify(string)
	receive(string)
}

type Participant1 struct {
	Mediator MediatorIf
	Name     string
	Id       int
}

func (p *Participant1) notify(msg string) {
	p.Mediator.SendMessage("Participant2", msg)
}

func (p *Participant1) receive(msg string) {
	fmt.Printf("%v received message: %v\n", p.Name, msg)
}

func (p *Participant1) ChangeID(id int) {
	msg := fmt.Sprintf("Participant1: changed id from %v to %v", p.Id, id)
	p.Id = id
	p.notify(msg)
}

type Participant2 struct {
	Mediator MediatorIf
	Name     string
	Data     string
}

func (p *Participant2) notify(msg string) {
	p.Mediator.SendMessage("Participant1", msg)
}

func (p *Participant2) receive(msg string) {
	fmt.Printf("%v received message: %v\n", p.Name, msg)
}

func (p *Participant2) ChangeData(data string) {
	msg := fmt.Sprintf("Participant2: changed data from %v to %v", p.Data, data)
	p.Data = data
	p.notify(msg)
}
