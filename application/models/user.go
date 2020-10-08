package models

type User struct {
	ID       uint64  `gorm:"primary_key" json:"id"`
	Name     string  `gorm:"size:255" json:"name"`
	Surname  string  `gorm:"size:255" json:"surname"`
	Password string  `gorm:"size:255" json:"-"`
	Email    string  `gorm:"size:255" json:"email"`
}

type UserAccount struct {
	User    *User
	Balance uint64
}
