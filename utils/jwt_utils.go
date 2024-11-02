package utils

import (
	"github.com/golang-jwt/jwt"
	"os"
	"time"
)

func GenerateJWT(email string) (string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	return accessToken.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
}
