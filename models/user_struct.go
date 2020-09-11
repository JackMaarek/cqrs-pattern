package models

type User struct {
	ID       uint64 `gorm:"primary_key"`
	Name     string `gorm:"size:255"`
	Surname  string `gorm:"size:255"`
	Password string `gorm:"size:255"`
	Email    string `gorm:"size:255"`
}
