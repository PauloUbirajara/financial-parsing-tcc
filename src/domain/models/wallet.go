package models

import (
	"errors"
	"time"
)

type Wallet struct {
	Id          string    `json:"id"`
	Currency    *Currency `json:"currency"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
}

func NewWallet(
	currency *Currency,
	name string,
	description string,
) (*Wallet, error) {
	if currency == nil {
		return nil, errors.New("Currency is null")
	}

	wallet := Wallet{
		Name:        name,
		Description: description,
		Currency:    currency,
		CreatedAt:   time.Now(),
	}
	return &wallet, nil
}
