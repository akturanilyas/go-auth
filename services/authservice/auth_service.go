package authservice

import (
	"go-test/models"
	"go-test/validations"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	*models.User
}

var db *gorm.DB

func SetDBConnection(conn *gorm.DB) {

	db = conn
}

func Register(user *validations.CreateUserValidation) *gorm.DB {
	record := models.User{Name: user.Name}

	result := db.Create(&record)

	return result
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
