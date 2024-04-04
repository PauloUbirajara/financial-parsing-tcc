package models

import (
	"time"

	"gorm.io/gorm"
)

type Wallet struct {
	ID        string `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Name        string `json:"name"`
	Description string `json:"description"`
}
