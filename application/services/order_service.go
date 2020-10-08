package services

import (
	"github.com/JackMaarek/cqrsPattern/application/structs/forms"
	"github.com/JackMaarek/cqrsPattern/application/validators"
	"github.com/JackMaarek/cqrsPattern/chore/cqrs"
	"github.com/JackMaarek/cqrsPattern/domain"
	"github.com/JackMaarek/cqrsPattern/domain/orders"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateOrder(c *gin.Context) (interface{}, error) {
	if err := validators.ValidateJsonHeader(c); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return nil, err
	}
	orderForm := forms.OrderForm{}
	if err := c.ShouldBindJSON(&orderForm); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return nil, err
	}
	command := cqrs.NewCommandMessage(&orders.CreateOrderCommand{OrderForm: &orderForm})
	o, err := domain.Cb.Dispatch(command)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Could not create user.")
		return nil, err
	}
	return o, nil
}
