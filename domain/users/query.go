package users

import (
	"errors"
	"github.com/JackMaarek/cqrsPattern/chore/cqrs"
)

type FindUsersQuery struct {
}

type FindUserByIdQuery struct {
	UserId uint64
}

type CreateUserQueryHandler struct{}

func (ch CreateUserQueryHandler) Handle(query cqrs.QueryMessage) (interface{}, error) {
	switch qry := query.Payload().(type) {
	case *FindUsersQuery:
		if ul, err := GetUsers(); err != nil {
			return nil, err
		} else {
			return ul, nil
		}
	case *FindUserByIdQuery:
		if u, err := GetUserByID(qry.UserId); err != nil {
			return nil, err
		} else {
			return u, nil
		}
	default:
		return nil, errors.New("bad query type")
	}
}

func NewCreateUserQueryHandler() *CreateUserQueryHandler {
	return &CreateUserQueryHandler{}
}
