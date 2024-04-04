package models

import (
	"time"

	"gorm.io/gorm"
)

type Currency_CurrencyRecord struct {
	ID        string `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	CurrencyId       string `json:"currency_id" gorm:"not null"`
	CurrencyRecordId string `json:"currency_record_id" gorm:"not null"`
}
