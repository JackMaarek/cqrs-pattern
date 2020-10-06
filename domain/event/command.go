package event

import (
	"errors"
	"github.com/JackMaarek/cqrsPattern/chore/cqrs"
	es "github.com/JackMaarek/cqrsPattern/chore/event-sourcing"
)

type CreateUserTopupAccountCommand struct {
	TopupEvent *es.TopUp
}

type CreateEventCommandHandler struct{}

func (ch CreateEventCommandHandler) Handle(command cqrs.CommandMessage) (interface{}, error) {
	switch cmd := command.Payload().(type) {
	case *CreateUserTopupAccountCommand:
		if err := CreateEventSnapshot(cmd.TopupEvent); err != nil {
			return nil, err
		}
		return nil, nil
	default:
		return nil, errors.New("bad query type")
	}
}

func NewCreateEventCommandHandler() *CreateEventCommandHandler {
	return &CreateEventCommandHandler{}
}
