package validators

import (
	"github.com/JackMaarek/cqrsPattern/application/models"
	"golang.org/x/crypto/bcrypt"
)

// Hash the password passed as argument
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// Verify the hashed password and the password to check password consistency
func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// Get the user model and hash it password before saving it
func BeforeSave(user *models.User) error {
	hashedPassword, err := Hash(user.Password)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return nil
}
