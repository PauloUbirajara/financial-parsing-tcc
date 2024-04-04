package models

import (
	"time"

	"gorm.io/gorm"
)

type Currency_User struct {
	ID        string `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	CurrencyId string `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserId     string `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
