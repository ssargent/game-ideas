package messaging

type GameMessage struct {
	Type    string
	Payload interface{}
}

type MessageService interface {
	Name() string
	Accept(g *GameMessage) bool
	Receive(msg *GameMessage) error
}

type GameMessageService interface {
	AddService(m MessageService)
	Send(g *GameMessage)
}

type gameMessageService struct {
	services map[string]MessageService
}

func (gms *gameMessageService) AddService(m MessageService) {
	if _, ok := gms.services[m.Name()]; !ok {
		gms.services[m.Name()] = m
	}
}

func (gms *gameMessageService) Send(g *GameMessage) {
	for _, v := range gms.services {
		if v.Accept(g) {
			v.Receive(g)
		}
	}
}
