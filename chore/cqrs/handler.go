package cqrs

type CommandHandler interface {
	Handle(CommandMessage) error
}

type QueryHandler interface {
	Handle(QueryMessage) (interface{}, error)
}
