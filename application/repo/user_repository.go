package repo

import (
	"errors"
	"fmt"
	"github.com/JackMaarek/cqrsPattern/application/models"
	"github.com/jinzhu/gorm"
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

// EditUserByID update a user from its Id.
func EditUserByID(user *models.User) (*models.User, error) {
	var err error
	err = models.DB.Debug().Save(&user).Take(&user).Error
	if err != nil {
		return &models.User{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &models.User{}, errors.New("User Not Found")
	}
	return user, err
}

// DeleteUserByID delete a user from its Id.
func DeleteUserByID(id uint64) (models.User, error) {
	var err error
	var user models.User

	err = models.DB.Debug().Delete(&user, id).Error
	if err != nil {
		return models.User{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return models.User{}, errors.New("User Not Found")
	}
	return user, err
}

// FindUsers returns you a list of all stored users.
func FindUsers() (*[]models.User, error) {
	var err error
	var users []models.User
	err = models.DB.Debug().Find(&users).Error
	if gorm.IsRecordNotFoundError(err) {
		fmt.Println("no record found")
		return &[]models.User{}, errors.New("Users Not Found")
	}
	return &users, err
}

// FindUsersById returns you a user from its Id.
func FindUserByID(uid uint64) (*models.User, error) {
	var err error
	var user models.User
	err = models.DB.Debug().Model(models.User{}).Where("id = ?", uid).Take(&user).Error
	if err != nil {
		return &models.User{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &models.User{}, errors.New("User Not Found")
	}
	return &user, err
}
