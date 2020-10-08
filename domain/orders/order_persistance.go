package orders

import (
	"github.com/JackMaarek/cqrsPattern/application/models"
	"github.com/JackMaarek/cqrsPattern/application/repo"
	"github.com/JackMaarek/cqrsPattern/application/structs/forms"
)

func PersistOrder(form *forms.OrderForm) (*models.Order, error) {
	u , err := repo.FindUserByID(1)
	if err != nil {
		return nil, err
	}
	o := models.Order{
		TotalPrice: form.TotalPrice,
		ItemCount:  form.ItemCount,
		UserId: u.ID,
	}
	co, err := repo.CreateOrder(&o)
	if err != nil {
		return nil, err
	}
	//err = es.NewOrderCreatedEvent(u)
	return co, nil
}