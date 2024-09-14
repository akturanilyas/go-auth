package auth

import (
	"auth-go/pkg/database"
	"errors"
	"time"

	"auth-go/pkg/user"
	"auth-go/utils"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	*user.User
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

func LoginService(loginData *LoginValidation) (string, error) {
	var user user.User
	if err := database.DB.Where("email = ?", loginData.Email).First(&user).Error; err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)); err != nil {
		return "", err
	}

	token, err := utils.GenerateJWT(user.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}

func RefreshTokenService(refreshToken string) (string, error) {
	claims := &utils.JWTClaim{}
	token, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		a, _ := utils.GetJWTSecret()
		return a, nil
	})

	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", errors.New("invalid refresh token")
	}

	if time.Unix(claims.ExpiresAt, 0).Before(time.Now()) {
		return "", errors.New("refresh token is expired")
	}

	var user user.User
	if err := database.DB.Where("email = ?", claims.Email).First(&user).Error; err != nil {
		return "", err
	}

	// Generate a new access token
	newAccessToken, err := utils.GenerateJWT(user.Email)
	if err != nil {
		return "", err
	}

	return newAccessToken, nil
}

func ForgotPasswordService(email string) error {
	resetToken := utils.GenerateResetToken()

	var user user.User
	if err := database.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return err
	}

	user.ResetToken = resetToken
	if err := database.DB.Save(&user).Error; err != nil {
		return err
	}
	// TODO: Send reset password email
	return nil
}
