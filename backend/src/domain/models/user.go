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

	Username     string `gorm:"uniqueIndex"`
	Email        string `gorm:"uniqueIndex"`
	PasswordHash string
	Active       bool
}
