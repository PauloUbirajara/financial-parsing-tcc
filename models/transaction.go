package models

import (
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

func NewTransaction(wallet *Wallet, name string, description string, value float64, transactionDate time.Time) *Transaction {
	transaction := Transaction{}
	return &transaction
}
