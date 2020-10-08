package repo

import (
	"errors"
	"fmt"
	"github.com/JackMaarek/cqrsPattern/application/conf"
	"github.com/JackMaarek/cqrsPattern/application/models"
	"github.com/jinzhu/gorm"
)

func CreateUser(user *models.User) (*models.User, error) {
	var err error
	err = conf.DB.Debug().Create(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func EditUserByID(id uint64, user *models.User) (*models.User, error) {
	var err error
	err = conf.DB.Debug().Model(&models.User{}).Where("id = ?", id).Update(&user).Take(&user).Error
	if err != nil {
		return nil, errors.New("Could'nt update user")
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("User Not Found")
	}
	return user, err
}

func DeleteUserByID(id uint64) error {
	var err error
	var user models.User

	err = conf.DB.Debug().Delete(&user, id).Error
	if err != nil {
		return err
	}
	if gorm.IsRecordNotFoundError(err) {
		return errors.New("User Not Found")
	}
	return nil
}

func FindUsers() (*[]models.User, error) {
	var err error
	var users []models.User
	err = conf.DB.Debug().Find(&users).Error
	if gorm.IsRecordNotFoundError(err) {
		fmt.Println("no record found")
		return nil, errors.New("Users Not Found")
	}
	return &users, err
}

func FindUserByID(uid uint64) (*models.User, error) {
	var err error
	var user models.User
	err = conf.DB.Debug().Model(models.User{}).Where("id = ?", uid).Take(&user).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("User Not Found")
	}
	return &user, err
}
