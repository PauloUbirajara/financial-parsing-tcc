package models

import (
	"time"

	"gorm.io/gorm"
)

type CurrencyRecord struct {
	ID        string `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Value      float64
	RecordDate time.Time
}
