package user

import (
	"github.com/JackMaarek/cqrsPattern/application/models"
	"github.com/JackMaarek/cqrsPattern/application/repo"
	"github.com/JackMaarek/cqrsPattern/application/structs/forms"
	"github.com/JackMaarek/cqrsPattern/application/validators"
)

func PersistUser(form *forms.UserForm) error {
	u := models.User{
		Name:     form.Name,
		Surname:  form.Surname,
		Password: form.Password,
		Email:    form.Email,
	}
	if err := validators.BeforeSave(&u); err != nil{
		return err
	}
	_, err := repo.CreateUser(&u)
	if err != nil {
		return err
	}
	return nil
}

func ModifyUser(form *forms.UserForm) error {
	u := models.User{
		Name:     form.Name,
		Surname:  form.Surname,
		Password: form.Password,
		Email:    form.Email,
	}
	_, err := repo.EditUserByID(&u)
	if err != nil {
		return err
	}
	return nil
}

func GetUsers() (*[]models.User, error) {
	ul, err := repo.FindUsers()
	println(ul)
	if err != nil {
		return nil, err
	}
	return ul, nil
}