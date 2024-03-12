package command

// паттерн Команда, позволяет инкапсулировать запрос как объект
// разделение на получатель действие

// Интерфейс, определяющий действие
type Command interface {
	execute()
}

// Структура, реализующая конкретную команду и содержащая приемник
type Presser struct {
	Device Receiver
}

func (p *Presser) execute() {
	p.Device.turnOnOff()
}

// Структура (Invoker), которая содержит объект команду и умеет ее запускать
type OnOffLight struct {
	Command Command
}

func (o *OnOffLight) TurnOnOffLight() {
	o.Command.execute()
}
