package user

import (
	"time"
)

type User struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	FirstName string     `gorm:"size:255;not null;" json:"first_name" validate:"min=3"`
	LastName  string     `gorm:"size:255;not null;" json:"last_name" validate:"min=3"`
	Email     string     `gorm:"size:255;unique;not null" json:"email" validate:"min=3"`
	Password  string     `gorm:"size:255;not null" json:"-" validate:"min=6"`
}
