package orders

import (
	"fmt"
	"github.com/JackMaarek/cqrsPattern/application/models"
	"github.com/JackMaarek/cqrsPattern/application/repo"
	"github.com/JackMaarek/cqrsPattern/application/structs/forms"
	"github.com/JackMaarek/cqrsPattern/chore/es"
)

func PersistOrder(form *forms.OrderForm) (*models.Order, error) {
	u, err := repo.FindUserByID(1)
	if err != nil {
		return nil, err
	}
	o := models.Order{
		TotalPrice: form.TotalPrice,
		ItemCount:  form.ItemCount,
		UserId:     u.ID,
	}
	co, err := repo.CreateOrder(&o)
	if err != nil {
		return nil, err
	}
	if err = es.NewOrderCreatedEvent(u); err != nil {
		fmt.Println(err)
	}
	return co, nil
}
