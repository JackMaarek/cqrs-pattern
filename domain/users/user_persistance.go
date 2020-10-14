package users

import (
	"fmt"
	"github.com/JackMaarek/cqrsPattern/application/models"
	"github.com/JackMaarek/cqrsPattern/application/repo"
	"github.com/JackMaarek/cqrsPattern/application/structs/forms"
	"github.com/JackMaarek/cqrsPattern/application/validators"
)

func GetUsers() (*[]models.User, error) {
	ul, err := repo.FindUsers()
	fmt.Println(ul)
	if err != nil {
		return nil, err
	}
	return ul, nil
}

func GetUserByID(id uint64) (*models.User, error) {
	u, err := repo.FindUserByID(id)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func PersistUser(form *forms.UserForm) (*models.User, error) {
	u := models.User{
		Name:     form.Name,
		Surname:  form.Surname,
		Password: form.Password,
		Email:    form.Email,
	}
	if err := validators.BeforeSave(&u); err != nil {
		return nil, err
	}
	usr, err := repo.CreateUser(&u)
	if err != nil {
		return nil, err
	}
	return usr, nil
}

func ModifyUser(id uint64, form *forms.UserForm) (*models.User, error) {
	u := models.User{
		Name:     form.Name,
		Surname:  form.Surname,
		Password: form.Password,
		Email:    form.Email,
	}
	if err := validators.BeforeSave(&u); err != nil {
		return nil, err
	}
	edtdusr, err := repo.EditUserByID(id, &u)
	if err != nil {
		return nil, err
	}
	return edtdusr, nil
}

func DeleteUser(id uint64) error {
	err := repo.DeleteUserByID(id)
	if err != nil {
		return err
	}
	return nil
}
