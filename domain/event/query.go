package event

import (
	"errors"
	"github.com/JackMaarek/cqrsPattern/chore/cqrs"
	es "github.com/JackMaarek/cqrsPattern/chore/event-sourcing"
)

type FindTopupEventQuery struct {
	Topup es.TopUp
}

type CreateEventQueryHandler struct{}

func (ch CreateEventQueryHandler) Handle(query cqrs.QueryMessage) (interface{}, error) {
	switch qry := query.Payload().(type) {
	case *FindTopupEventQuery:
		if err := RetrieveEventSnapshot(&qry.Topup); err != nil {
			return nil, err
		} else {
			return nil, nil
		}

	default:
		return nil, errors.New("bad query type")
	}
}

func NewCreateTopupQueryHandler() *CreateEventQueryHandler {
	return &CreateEventQueryHandler{}
}
