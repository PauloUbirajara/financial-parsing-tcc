package models

import (
	"errors"
	"time"
)

type Wallet struct {
	Id string

	Currency *Currency

	Name        string
	Description string

	CreatedAt time.Time
}

func NewWallet(
	currency *Currency,
	name string,
	description string,
) (*Wallet, error) {
	if currency == nil {
		return nil, errors.New("Currency is null")
	}

	wallet := Wallet{}
	return &wallet, nil
}
