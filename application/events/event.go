package events

import (
	"fmt"
)

const TopUpType Type = "TopUp"

type Type string

type Event struct {
	ID   string
	Type Type
}

func (o *Event) GetID() string {
	return o.ID
}

func (o *Event) SetID(id string) {
	o.ID = id
}

func (o *Event) GetType() Type {
	return o.Type
}
func (o *Event) String() string {

	return fmt.Sprintf("id:%s type:%s", o.ID, o.Type)
}

type EventInterface interface {
	GetID() string
	GetType() Type
	SetID(id string)
}
