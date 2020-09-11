package repo

import (
	"github.com/JackMaarek/cqrsPattern/models"
)

// CreateUser creates an user row in database
func CreateUser(user *models.User) (*models.User, error) {
	var err error
	err = models.DB.Debug().Create(&user).Error
	if err != nil {
		return &models.User{}, err
	}
	return user, nil
}
