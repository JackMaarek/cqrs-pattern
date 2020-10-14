package orders

import (
	"errors"
	"github.com/JackMaarek/cqrsPattern/application/structs/forms"
	"github.com/JackMaarek/cqrsPattern/chore/cqrs"
)

type CreateOrderCommand struct {
	OrderForm *forms.OrderForm
}

type CreateOrderCommandHandler struct{}

func (och CreateOrderCommandHandler) Handle(command cqrs.CommandMessage) (interface{}, error) {
	switch cmd := command.Payload().(type) {
	case *CreateOrderCommand:
		if o, err := PersistOrder(cmd.OrderForm); err != nil {
			return nil, err
		} else {
			return o, err
		}
	default:
		return nil, errors.New("unknown order command type")
	}
}

func NewCreateOrderCommandHandler() *CreateOrderCommandHandler {
	return &CreateOrderCommandHandler{}
}
