package models

import (
	"errors"
	"time"
)

type Transaction struct {
	Id                string    `form:"id" json:"id"`
	TransactionTypeId string    `form:"TransactionTypeId" json:"TransactionTypeId"`
	Wallet            *Wallet   `form:"wallet" json:"wallet"`
	Name              string    `form:"name" json:"name"`
	Description       string    `form:"description" json:"description"`
	Value             float64   `form:"value" json:"value"`
	TransactionDate   time.Time `form:"transactionDate" json:"transactionDate"`
	CreatedAt         time.Time `form:"createdAt" json:"createdAt"`
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
		Name:            name,
		Description:     description,
		TransactionDate: transactionDate,
		Value:           value,
		Wallet:          wallet,
		CreatedAt:       time.Now(),
	}
	return &transaction, nil
}
