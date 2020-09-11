package handlers

import (
	"fmt"
	"github.com/JackMaarek/cqrsPattern/models"
	"github.com/JackMaarek/cqrsPattern/repo"
	"github.com/JackMaarek/cqrsPattern/structs"
	"github.com/JackMaarek/cqrsPattern/validators"
)

func CreateUserCommand(c *structs.Command) (*models.User, error) {
	userForm := c.Data.(structs.UserForm)
	u := models.User{
		Name:     userForm.Name,
		Surname:  userForm.Surname,
		Password: userForm.Password,
		Email:    userForm.Email,
	}
	err := validators.BeforeSave(&u)
	if err != nil {
		return nil, err
	}
	fmt.Println(u)
	cu, err := repo.CreateUser(&u)
	if err != nil {
		return nil, err
	}
	return cu, nil
}
