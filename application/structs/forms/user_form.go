package forms

type UserForm struct {
	Name     string `json:"name" validate:"required"`
	Surname  string `json:"surname" validate:"required"`
	Password string `json:"password" validate:"required,password"`
	Email    string `json:"email" validate:"required,email"`
}
