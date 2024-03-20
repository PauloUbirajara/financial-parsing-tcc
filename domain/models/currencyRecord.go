package models

import (
	"errors"
	"time"
)

type CurrencyRecord struct {
	Id string

	Currency *Currency

	Value      float64
	RecordDate time.Time

	CreatedAt time.Time
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
		Value: value,
		RecordDate: recordDate,

		Currency: currency,

		CreatedAt: time.Now(),
	}
	return &currencyRecord, nil
}
