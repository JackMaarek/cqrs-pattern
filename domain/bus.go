package domain

import (
	"github.com/JackMaarek/cqrsPattern/chore/cqrs"
	"github.com/JackMaarek/cqrsPattern/domain/user"
)

var Cb *cqrs.CommandBus
var Qb *cqrs.QueryBus

func InitBuses()  {
	Cb = cqrs.NewCommandBus()
	Qb = cqrs.NewQueryBus()

	_ = Cb.RegisterHandler(user.NewCreateUserCommandHandler(), &user.CreateUserCommand{})
	_ = Cb.RegisterHandler(user.NewCreateUserCommandHandler(), &user.PUTUserCommand{})
	_ = Qb.RegisterHandler(user.NewCreateUserQueryHandler(), &user.FindUserQuery{})
}