package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        string `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Email        string
	PasswordHash string
	Active       bool
}
