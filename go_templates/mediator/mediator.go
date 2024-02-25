package mediator

type MediatorIf interface {
	SendMessage(participant string, msg string)
}

type ConcreteMediator struct {
	participants map[string]ParticipantIf
}

func (cm *ConcreteMediator) SendMessage(participant string, msg string) {
	cm.participants[participant].receive(msg)
}

func (cm *ConcreteMediator) RegisterParticipants(participants ...ParticipantIf) {
	cm.participants = make(map[string]ParticipantIf)
	for _, participant := range participants {
		if p, ok := participant.(*Participant1); ok {
			cm.participants[p.Name] = participant
		}
		if p, ok := participant.(*Participant2); ok {
			cm.participants[p.Name] = participant
		}
	}
}
