package models

import (
	"time"
)

type CurrencyRecord struct {
	Id         string    `json:"id"`
	Currency   *Currency `json:"currency"`
	Value      float64   `json:"value"`
	RecordDate time.Time `json:"recordDate"`
	CreatedAt  time.Time `json:"createdAt"`
}
