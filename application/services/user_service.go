package services

import (
	"github.com/JackMaarek/cqrsPattern/application/structs"
	"github.com/JackMaarek/cqrsPattern/application/structs/forms"
	"github.com/JackMaarek/cqrsPattern/chore/cqrs"
	"github.com/JackMaarek/cqrsPattern/domain"
	"github.com/JackMaarek/cqrsPattern/domain/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(c *gin.Context) (*user.CreateUserCommand, error) {
	header := structs.Header{}
	if err := c.ShouldBindHeader(&header); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		c.Abort()
	}
	if header.ContentType == "application/json" {
		c.JSON(http.StatusBadRequest, "Wrong content-type header")
		c.Abort()
	}
	userForm := forms.UserForm{}
	if err := c.ShouldBindJSON(&userForm); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		c.Abort()
	}
	command := cqrs.NewCommandMessage(&user.CreateUserCommand{UserForm: &userForm})
	_ = domain.Cb.Dispatch(command)

	return nil, nil
}

func GetUserList(c *gin.Context) (interface{}, error) {
	header := structs.Header{}
	if err := c.ShouldBindHeader(&header); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		c.Abort()
	}
	if header.ContentType == "application/json" {
		c.JSON(http.StatusBadRequest, "Wrong content-type header")
		c.Abort()
	}
	query := cqrs.NewQueryMessage(&user.FindUserQuery{})
	i, err := domain.Qb.Dispatch(query)
	if err != nil {
		c.JSON(http.StatusNoContent, err.Error())
		c.Abort()
	}
	return i, nil
}
