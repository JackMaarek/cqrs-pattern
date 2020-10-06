package services

import (
	"errors"
	"github.com/JackMaarek/cqrsPattern/application/structs"
	"github.com/JackMaarek/cqrsPattern/application/structs/forms"
	"github.com/JackMaarek/cqrsPattern/application/util"
	"github.com/JackMaarek/cqrsPattern/chore/cqrs"
	"github.com/JackMaarek/cqrsPattern/domain"
	"github.com/JackMaarek/cqrsPattern/domain/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(c *gin.Context) (interface{}, error) {
	header := structs.Header{}
	if err := c.ShouldBindHeader(&header); err != nil {
		return nil, err
	}
	if header.ContentType != "application/json" {
		return nil, errors.New("Wrong content type.")
	}
	userForm := forms.UserForm{}
	if err := c.ShouldBindJSON(&userForm); err != nil {
		return nil, err
	}
	command := cqrs.NewCommandMessage(&user.CreateUserCommand{UserForm: &userForm})
	usr, err := domain.Cb.Dispatch(command)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return nil, err
	}
	return usr, nil
}

func UpdateUser(c *gin.Context) (interface{}, error){
	header := structs.Header{}
	if err := c.ShouldBindHeader(&header); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		c.Abort()
	}
	if header.ContentType != "application/json" {
		c.JSON(http.StatusBadRequest, "Wrong content-type header")
		c.Abort()
	}
	userForm := forms.UserForm{}
	if err := c.ShouldBindJSON(&userForm); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		c.Abort()
	}
	id := util.ParseStringToUint64(c.Param("id"))
	command := cqrs.NewCommandMessage(&user.PUTUserCommand{UserId: id, UserForm: &userForm})
	usr, err := domain.Cb.Dispatch(command)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return nil, err
	}
	return usr, nil
}

func GetUserList(c *gin.Context) (interface{}, error) {
	query := cqrs.NewQueryMessage(&user.FindUserQuery{})
	i, err := domain.Qb.Dispatch(query)
	if err != nil {
		c.JSON(http.StatusNoContent, "")
		return nil, err
	}
	return i, nil
}
