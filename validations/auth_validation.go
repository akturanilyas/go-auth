package validations

type CreateUserValidation struct {
	AbstractValidation

	Name     string `validate:"required,min=5,max=20"`
	Password string `validate:"min=8,max=20"`
}
