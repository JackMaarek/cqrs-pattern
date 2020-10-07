package es

import (
	"encoding/json"
	"fmt"
)

type Type string

type Event struct {
	ID   uint64
	Type Type
	Data interface{}
}

func (e *Event) GetID() uint64 {
	return e.ID
}

func (e *Event) SetID(id uint64) {
	e.ID = id
}

func (e *Event) GetType() Type {
	return e.Type
}

func (e *Event) SetType(etype Type) {
	e.Type = etype
}

func (e *Event) GetData() interface{}  {
	return e.Data
}

func (e *Event) SetData(d interface{})  {
	e.Data = d
}

func (e *Event) String() string {

	return fmt.Sprintf("id:%s type:%s data:%s", e.ID, e.Type, e.Data)
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
	GetID() uint64
	GetType() Type
	SetID(id uint64)
	SetType(etype Type)
	GetData() interface{}
	SetData(d interface{})
}
