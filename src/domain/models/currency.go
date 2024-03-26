package models

import (
	"time"

	"gorm.io/gorm"
)

type Currency struct {
	ID        string `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Name           string
	Representation *string
}
