package models

import (
	"errors"
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

func NewCurrencyRecord(
	currency *Currency,
	name string,
	value float64,
	recordDate time.Time,
) (*CurrencyRecord, error) {
	if currency == nil {
		return nil, errors.New("Currency is null")
	}

	currencyRecord := CurrencyRecord{}
	return &currencyRecord, nil
}
