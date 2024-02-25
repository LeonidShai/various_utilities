package mediator

import "fmt"

type ParticipantIf interface {
	notify(string, string)
	receive(string)
}

type Participant struct {
	Mediator MediatorIf
	Name     string
}

func (p *Participant) notify(dst string, msg string) {
	fmt.Println("General notify")
	p.Mediator.SendMessage(dst, msg)
}

func (p *Participant) receive(msg string) {
	fmt.Printf("%v received message: %v\n", p.Name, msg)
}

type Participant1 struct {
	Participant

	Id int
}

func (p *Participant1) notify(dst string, msg string) {
	fmt.Println("Participant1 notify")
	p.Mediator.SendMessage(dst, msg)
}

func (p *Participant1) ChangeID(id int) {
	msg := fmt.Sprintf("Participant1 changed id from %v to %v", p.Id, id)
	p.Id = id
	p.notify("Participant2", msg)
}

type Participant2 struct {
	Participant

	Data string
}

func (p *Participant2) ChangeData(data string) {
	msg := fmt.Sprintf("Participant2 changed data from %v to %v", p.Data, data)
	p.Data = data
	p.notify("Participant1", msg)
}
