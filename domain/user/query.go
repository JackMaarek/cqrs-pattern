package user

import (
	"errors"
	"fmt"
	"github.com/JackMaarek/cqrsPattern/chore/cqrs"
)

type FindUserQuery struct {

}

type CreateUserQueryHandler struct {}

func (ch CreateUserQueryHandler) Handle(query cqrs.QueryMessage) (interface{}, error) {
	switch qry := query.Payload().(type) {
	case *FindUserQuery:
		fmt.Println(qry)
		ul, err := GetUsers()
		fmt.Println(ul)
		if err != nil {
			return nil, err
		}
		return ul, nil
	default:
		return nil, errors.New("bad query type")
	}
}

func NewCreateUserQueryHandler() *CreateUserQueryHandler {
	return &CreateUserQueryHandler{}
}
