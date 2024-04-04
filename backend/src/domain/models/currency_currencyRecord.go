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

	CurrencyId       string `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" gorm:"foreignKey: CurrencyID"`
	CurrencyRecordId string `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" gorm:"foreignKey: CurrencyRecordID"`
}
