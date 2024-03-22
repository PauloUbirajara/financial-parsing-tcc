package models

import (
	"errors"
	"time"
)

type CurrencyRecord struct {
	Id         string    `json:"id"`
	Currency   *Currency `json:"currency"`
	Value      float64   `json:"value"`
	RecordDate time.Time `json:"recordDate"`
	CreatedAt  time.Time `json:"createdAt"`
}

func NewCurrencyRecord(
	currency *Currency,
	value float64,
	recordDate time.Time,
) (*CurrencyRecord, error) {
	if currency == nil {
		return nil, errors.New("Currency is null")
	}

	currencyRecord := CurrencyRecord{
		Value:      value,
		RecordDate: recordDate,
		Currency:   currency,
		CreatedAt:  time.Now(),
	}
	return &currencyRecord, nil
}
