package models

import (
	"time"
)

type CurrencyRecord struct {
	Id string

	Currency *Currency

	Name       string
	Value      float64
	RecordDate time.Time

	CreatedAt time.Time
}
