package models

import (
	"errors"
	"time"
)

type Transaction struct {
	Id string

	TransactionTypeId string
	Wallet            *Wallet

	Name            string
	Description     string
	Value           float64
	TransactionDate time.Time

	CreatedAt time.Time
}

func NewTransaction(
	wallet *Wallet,
	name string,
	description string,
	value float64,
	transactionDate time.Time,
) (*Transaction, error) {
	if wallet == nil {
		return nil, errors.New("Wallet is null")
	}

	transaction := Transaction{
		Name: name,
		Description: description,
		TransactionDate: transactionDate,
		Value: value,

		Wallet: wallet,

		CreatedAt: time.Now(),
	}
	return &transaction, nil
}
