package domain

import (
	"github.com/JackMaarek/cqrsPattern/chore/cqrs"
	"github.com/JackMaarek/cqrsPattern/domain/orders"
	"github.com/JackMaarek/cqrsPattern/domain/users"
)

var Cb *cqrs.CommandBus
var Qb *cqrs.QueryBus

func InitBuses() {
	Cb = cqrs.NewCommandBus()
	Qb = cqrs.NewQueryBus()

	//>------------------ User Commands --------------------------
	_ = Cb.RegisterHandler(users.NewCreateUserCommandHandler(), &users.CreateUserCommand{})
	_ = Cb.RegisterHandler(users.NewCreateUserCommandHandler(), &users.PUTUserCommand{})
	_ = Cb.RegisterHandler(users.NewCreateUserCommandHandler(), &users.DeleteUserCommand{})
	//<------------------ End User Commands --------------------------

	//>------------------ User Queries --------------------------
	_ = Qb.RegisterHandler(users.NewCreateUserQueryHandler(), &users.FindUsersQuery{})
	_ = Qb.RegisterHandler(users.NewCreateUserQueryHandler(), &users.FindUserByIdQuery{})
	//<------------------ End User Queries --------------------------


	//>------------------ Order Commands --------------------------
	_ = Cb.RegisterHandler(orders.NewCreateOrderCommandHandler(), &orders.CreateOrderCommand{})
	//<------------------ End Order Commands --------------------------

}
