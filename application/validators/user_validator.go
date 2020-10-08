package validators

import (
	"errors"
	"github.com/JackMaarek/cqrsPattern/application/models"
	"github.com/JackMaarek/cqrsPattern/application/structs/forms"
	"github.com/badoux/checkmail"
	"golang.org/x/crypto/bcrypt"
	"strings"
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

func ValidateUser(user *forms.UserForm, action string) error {
	switch strings.ToLower(action) {
	case "update":
		if user.Name == "" {
			return errors.New("Required Name")
		}
		if user.Password == "" {
			return errors.New("Required Password")
		}
		if user.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(user.Email); err != nil {
			return errors.New("Invalid Email")
		}

		return nil
	case "login":
		if user.Password == "" {
			return errors.New("Required Password")
		}
		if user.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(user.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil

	default:
		if user.Name == "" {
			return errors.New("Require Name")
		}
		if user.Email == "" {
			return errors.New("Required Email")
		}
		if user.Password == "" {
			return errors.New("Required Password")
		}
		if err := checkmail.ValidateFormat(user.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil
	}
}
