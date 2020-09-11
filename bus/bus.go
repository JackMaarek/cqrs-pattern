package bus

import (
	"fmt"
	"github.com/JackMaarek/cqrsPattern/handlers"
	"github.com/JackMaarek/cqrsPattern/structs"
)

func HandleRequest(c *structs.Command) (*structs.Command, error) {
	fmt.Println("handlerRequest triggered")
	switch c.Type {
	case "User Creation Command":
		u, err := handlers.CreateUserCommand(c)
		if err != nil {
			return nil, err
		}
		c.Data = &u
		return c, nil
	default:
		c.Data = ""
		return c, nil
	}
	//data := channels.Command{}
}
