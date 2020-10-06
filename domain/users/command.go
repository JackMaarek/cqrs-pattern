package users

import (
	"errors"
	"github.com/JackMaarek/cqrsPattern/application/structs/forms"
	"github.com/JackMaarek/cqrsPattern/chore/cqrs"
)

type CreateUserCommand struct {
	UserForm *forms.UserForm
}

type PUTUserCommand struct {
	UserId   uint64
	UserForm *forms.UserForm
}

type DeleteUserCommand struct {
	UserId uint64
}

type CreateUserCommandHandler struct{}

func (ch CreateUserCommandHandler) Handle(command cqrs.CommandMessage) (interface{}, error) {

	switch cmd := command.Payload().(type) {
	case *CreateUserCommand:
		if usr, err := PersistUser(cmd.UserForm); err != nil {
			return nil, err
		} else {
			return usr, nil
		}
	case *PUTUserCommand:
		if usr, err := ModifyUser(cmd.UserId, cmd.UserForm); err != nil {
			return nil, err
		} else {
			return usr, nil
		}
	case *DeleteUserCommand:
		if err := DeleteUser(cmd.UserId); err != nil {
			return nil, err
		}
		return nil, nil
	default:
		return nil, errors.New("bad command type")
	}
}

func NewCreateUserCommandHandler() *CreateUserCommandHandler {
	return &CreateUserCommandHandler{}
}
