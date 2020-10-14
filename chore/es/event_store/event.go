package estore

import (
	"encoding"
	"encoding/json"
	"fmt"
)

type Type string

type Event struct {
	ID   string
	Type Type
}

func (e *Event) GetID() string {
	return e.ID
}

func (e *Event) SetID(id string) {
	e.ID = id
}

func (e *Event) GetType() Type {
	return e.Type
}

func (e *Event) SetType(etype Type) {
	e.Type = etype
}

func (e *Event) String() string {

	return fmt.Sprintf("id:%s type:%s", e.ID, e.Type)
}

func (e *Event) MarshalBinary() ([]byte, error) {
	return json.Marshal(e)
}

func (e *Event) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &e); err != nil {
		return err
	}

	return nil
}

type EventInterface interface {
	GetID() string
	GetType() Type
	SetID(id string)
	SetType(etype Type)
	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler
}
