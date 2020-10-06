package domain

import (
	"github.com/JackMaarek/cqrsPattern/chore/cqrs"
	"github.com/JackMaarek/cqrsPattern/domain/users"
)

var Cb *cqrs.CommandBus
var Qb *cqrs.QueryBus

func InitBuses()  {
	Cb = cqrs.NewCommandBus()
	Qb = cqrs.NewQueryBus()

	_ = Cb.RegisterHandler(users.NewCreateUserCommandHandler(), &users.CreateUserCommand{})
	_ = Cb.RegisterHandler(users.NewCreateUserCommandHandler(), &users.PUTUserCommand{})
	_ = Cb.RegisterHandler(users.NewCreateUserCommandHandler(), &users.DeleteUserCommand{})
	_ = Qb.RegisterHandler(users.NewCreateUserQueryHandler(), &users.FindUsersQuery{})
	_ = Qb.RegisterHandler(users.NewCreateUserQueryHandler(), &users.FindUserByIdQuery{})
}