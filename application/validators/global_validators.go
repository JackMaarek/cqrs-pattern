package validators

import (
	"github.com/badoux/checkmail"
	"github.com/go-playground/validator/v10"
	"regexp"
)

var V = validator.New()

func InitValidator() {

	_ = V.RegisterValidation("password", func(fl validator.FieldLevel) bool {
		r, _ := regexp.MatchString("^(?=.*?[A-Z])(?=.*?[a-z])(?=.*?[0-9])(?=.*?[#?!@$%^&*-]).{8,}$", fl.Field().String())
		return r
	})

	_ = V.RegisterValidation("email", func(fl validator.FieldLevel) bool {
		err := checkmail.ValidateFormat(fl.Field().String())
		if err != nil {
			return false
		}
		return true
	})
}
