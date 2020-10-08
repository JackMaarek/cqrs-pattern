package conf

import "github.com/JackMaarek/cqrsPattern/application/models"

// MakeMigrations executes all migrations for our structs
func MakeMigrations() {
	DB.AutoMigrate(&models.User{}, &models.Order{})
}
