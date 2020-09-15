package models

// MakeMigrations executes all migrations for our structs
func MakeMigrations() {
	DB.AutoMigrate(&User{})
}
