package auth

type CreateUserValidation struct {
	FirstName string `form:"first_name" validate:"required,min=3,max=20"`
	LastName  string `form:"last_name" validate:"required,min=3,max=20"`
	Email     string `form:"email" validate:"required,email,min=3,max=50"`
	Password  string `form:"password" validate:"required,min=6,max=20"`
}

type LoginValidation struct {
	Email    string `form:"email" validate:"required,email,min=3,max=50"`
	Password string `form:"password" validate:"required,min=6,max=20"`
}
