package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID        string `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	TransactionTypeId string
	WalletId          string
	Name              string
	Description       string
	Value             float64
	TransactionDate   time.Time
}
