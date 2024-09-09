package auth

import (
	"go-test/pkg/user"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	*user.User
}

var db *gorm.DB

func SetDBConnection(conn *gorm.DB) {
	db = conn
}

func RegisterService(_userData *CreateUserValidation) *gorm.DB {
	record := user.User{Name: _userData.Name}

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
