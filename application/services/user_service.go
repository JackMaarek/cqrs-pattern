package services

import (
	"github.com/JackMaarek/cqrsPattern/application/structs/forms"
	"github.com/JackMaarek/cqrsPattern/application/util"
	"github.com/JackMaarek/cqrsPattern/application/validators"
	"github.com/JackMaarek/cqrsPattern/chore/cqrs"
	"github.com/JackMaarek/cqrsPattern/domain"
	"github.com/JackMaarek/cqrsPattern/domain/users"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserList(c *gin.Context) (interface{}, error) {
	query := cqrs.NewQueryMessage(&users.FindUsersQuery{})
	ul, err := domain.Qb.Dispatch(query)
	if err != nil {
		c.JSON(http.StatusNoContent, "")
		return nil, err
	}
	return ul, nil
}

func GetUser(c *gin.Context) (interface{}, error) {
	id := util.ParseStringToUint64(c.Param("id"))
	query := cqrs.NewQueryMessage(&users.FindUserByIdQuery{UserId: id})
	u, err := domain.Qb.Dispatch(query)
	if err != nil {
		c.JSON(http.StatusNoContent, "")
		return nil, err
	}
	return u, nil
}

func CreateUser(c *gin.Context) (interface{}, error) {
	if err := validators.ValidateJsonHeader(c); err != nil {
		return nil, err
	}
	userForm := forms.UserForm{}
	if err := c.ShouldBindJSON(&userForm); err != nil {
		return nil, err
	}
	command := cqrs.NewCommandMessage(&users.CreateUserCommand{UserForm: &userForm})
	usr, err := domain.Cb.Dispatch(command)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Could not create user.")
		return nil, err
	}
	return usr, nil
}

func UpdateUser(c *gin.Context) (interface{}, error) {
	if err := validators.ValidateJsonHeader(c); err != nil {
		return nil, err
	}
	userForm := forms.UserForm{}
	if err := c.ShouldBindJSON(&userForm); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return nil, err
	}
	id := util.ParseStringToUint64(c.Param("id"))
	command := cqrs.NewCommandMessage(&users.PUTUserCommand{UserId: id, UserForm: &userForm})
	usr, err := domain.Cb.Dispatch(command)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return nil, err
	}
	return usr, nil
}

func RemoveUser(c *gin.Context) error {
	id := util.ParseStringToUint64(c.Param("id"))
	command := cqrs.NewCommandMessage(&users.DeleteUserCommand{UserId: id})
	_, err := domain.Cb.Dispatch(command)
	if err != nil {
		c.JSON(http.StatusNotModified, "")
		return err
	}
	return nil
}
