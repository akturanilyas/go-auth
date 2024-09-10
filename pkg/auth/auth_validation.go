package auth

type CreateUserValidation struct {
	FirstName string `validate:"required,min=3,max=20"`
	LastName  string `validate:"required,min=3,max=20"`
	Email     string `validate:"required,min=3,max=20"`
	Password  string `validate:"min=6,max=20"`
}
