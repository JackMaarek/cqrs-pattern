package event

import (
	"errors"
	"github.com/JackMaarek/cqrsPattern/application/events"
	"github.com/JackMaarek/cqrsPattern/chore/cqrs"
)

type FindTopupEventQuery struct {
	Topup events.TopUp
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
