package es

import (
	"fmt"
	"github.com/JackMaarek/cqrsPattern/chore/es/event_store"
)

const (
	UserCreated  estore.Type = "User Created"
	UserUpdated  estore.Type = "User Updated"
	OrderCreated estore.Type = "Order Created"
)

func NewEvent(t estore.Type) (*estore.Event, error) {
	e := &estore.Event{
		Type: t,
	}
	switch t {
	case OrderCreated:
		return e, nil
	default:
		return nil, fmt.Errorf("cannot find type %v", t)
	}
}

func NewTypedEvent(t estore.Type) (estore.EventInterface, error) {
	e := &estore.Event{
		Type: t,
	}
	switch t {
	case OrderCreated:
		return &estore.OrderEvent{
			Event: e,
		}, nil
	default:
		return nil, fmt.Errorf("cannot find type %v", t)
	}
}
