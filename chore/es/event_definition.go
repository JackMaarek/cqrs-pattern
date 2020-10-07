package es

import (
	"fmt"
)

const (
	UserCreated Type = "User Created"
	UserUpdated Type = "User Updated"
)

func NewEvent(t Type) (*Event, error)   {
	e := &Event{
		Type: t,
	}
	switch t {
	case UserCreated:
		return e, nil
	case UserUpdated:
		return e, nil
	default:
		return nil, fmt.Errorf("cannot find type %v", t)
	}

}