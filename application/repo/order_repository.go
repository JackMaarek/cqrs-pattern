package repo

import (
	"errors"
	"fmt"
	"github.com/JackMaarek/cqrsPattern/application/conf"
	"github.com/JackMaarek/cqrsPattern/application/models"
	"github.com/jinzhu/gorm"
)

func CreateOrder(order *models.Order) (*models.Order, error) {
	var err error
	err = conf.DB.Debug().Create(&order).Error
	if err != nil {
		return nil, err
	}
	return order, nil
}

func EditOrderByID(id uint64, order *models.Order) (*models.Order, error) {
	var err error
	err = conf.DB.Debug().Model(&models.Order{}).Where("id = ?", id).Update(&order).Take(&order).Error
	if err != nil {
		return nil, errors.New("Could'nt update order")
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("Order Not Found")
	}
	return order, err
}

func DeleteOrderByID(id uint64) error {
	var err error
	var order models.Order

	err = conf.DB.Debug().Delete(&order, id).Error
	if err != nil {
		return err
	}
	if gorm.IsRecordNotFoundError(err) {
		return errors.New("Order Not Found")
	}
	return nil
}

func FindOrders() (*[]models.Order, error) {
	var err error
	var orders []models.Order
	err = conf.DB.Debug().Find(&orders).Error
	if gorm.IsRecordNotFoundError(err) {
		fmt.Println("no record found")
		return nil, errors.New("Orders Not Found")
	}
	return &orders, err
}

func FindOrderByID(uid uint64) (*models.Order, error) {
	var err error
	var order models.Order
	err = conf.DB.Debug().Model(models.Order{}).Where("id = ?", uid).Take(&order).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("Order Not Found")
	}
	return &order, err
}
