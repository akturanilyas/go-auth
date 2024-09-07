package models

import (
	"go-test/enums"
	"time"
)

type User struct {
	ID        uint         `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	Status    enums.Status `gorm:"size:20;not null;default:active" json:"status"`
	Name      string       `gorm:"size:255;not null" json:"name"`
	Email     string       `gorm:"size:255;unique;not null" json:"email"`
	Password  string       `gorm:"size:255;not null" json:"-"`
}
