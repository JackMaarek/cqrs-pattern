package user

import (
	"errors"
	"fmt"
	"github.com/JackMaarek/cqrsPattern/chore/cqrs"
)

type FindUserQuery struct {

}

type CreateUserQueryHandler struct {}

func (ch CreateUserQueryHandler) Handle(command cqrs.QueryMessage) (interface{}, error) {
	fmt.Println(command)
	switch cmd := command.Payload().(type) {
	case *FindUserQuery:
		fmt.Println(cmd)
		ul, err := GetUsers()
		if err != nil {
			return nil, err
		}
		return ul, nil
	default:
		return nil, errors.New("bad command type")
	}
}

func NewCreateUserQueryHandler() *CreateUserQueryHandler {
	return &CreateUserQueryHandler{}
}
