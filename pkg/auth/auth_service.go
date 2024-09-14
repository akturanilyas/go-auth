package auth

import (
	"auth-go/pkg/database"
	"auth-go/pkg/user"
	"auth-go/utils"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	user.User
}

func RegisterService(userData *CreateUserValidation) (*user.User, error) {
	newUser := &user.User{
		FirstName: userData.FirstName,
		LastName:  userData.LastName,
		Email:     userData.Email,
	}

	if err := newUser.HashPassword(userData.Password); err != nil {
		return nil, err
	}

	result := database.DB.Create(newUser)
	if result.Error != nil {
		return nil, result.Error
	}

	return newUser, nil
}

func LoginService(loginData *LoginValidation) (string, error) {
	var _user User

	if err := database.DB.Select("email, password").Where("email = ?", loginData.Email).First(&_user).Error; err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(_user.Password), []byte(loginData.Password)); err != nil {
		return "", err
	}

	token, err := utils.GenerateJWT(_user.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}

func RefreshTokenService(refreshToken string) {

}

func ForgotPasswordService(email string) {

}

func (u *User) hashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	u.Password = string(bytes)
	return nil
}

func (u *User) checkPassword(providedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(providedPassword))
}
