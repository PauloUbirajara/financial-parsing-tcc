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

	CurrencyId       string `gorm:"foreignKey:CompanyRefer" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE`
	CurrencyRecordId string `gorm:"foreignKey:CompanyRefer" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE`
}
