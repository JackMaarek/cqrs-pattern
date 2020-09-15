package user

import (
	"errors"
	"github.com/JackMaarek/cqrsPattern/application/structs/forms"
	"github.com/JackMaarek/cqrsPattern/chore/cqrs"
)

type CreateUserCommand struct {
	UserForm *forms.UserForm
}

type PUTUserCommand struct {
	UserForm *forms.UserForm
}

type DeleteUserCommand struct {
	ID uint64
}
type CreateUserCommandHandler struct {}

func (ch CreateUserCommandHandler) Handle(command cqrs.CommandMessage) error {

	switch cmd := command.Payload().(type) {
	case *CreateUserCommand:
		if err := PersistUser(cmd.UserForm); err != nil {
			return err
		}
	case *PUTUserCommand:
		if err := ModifyUser(cmd.UserForm); err != nil {
			return nil
		}
	default:
		return errors.New("bad command type")
	}

	return nil
}

func NewCreateUserCommandHandler() *CreateUserCommandHandler {
	return &CreateUserCommandHandler{}
}
