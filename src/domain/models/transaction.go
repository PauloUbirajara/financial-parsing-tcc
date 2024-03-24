package models

import (
	"time"
)

type CreateTransactionInput struct {
	TransactionTypeId string    `form:"TransactionTypeId" json:"TransactionTypeId"`
	WalletId          string    `form:"wallet" json:"wallet"`
	Name              string    `form:"name" json:"name"`
	Description       string    `form:"description" json:"description"`
	Value             float64   `form:"value" json:"value"`
	TransactionDate   time.Time `form:"transactionDate" json:"transactionDate"`
}

type CreateTransactionOutput struct {
	Id    string `form:"id" json:"id"`
	Error error  `form:"error" json:"error"`
}

type GetTransactionByIdInput struct {
	Id string `form:"id" json:"id"`
}

type GetTransactionByIdOutput struct {
	Transaction Transaction
}

type GetAllTransactionsInput struct{}

type GetAllTransactionsOutput struct {
	Transactions []Transaction
	Error        error `form:"error" json:"error"`
}

type UpdateTransactionInput struct {
	Id                string    `form:"id" json:"id"`
	TransactionTypeId string    `form:"TransactionTypeId" json:"TransactionTypeId"`
	WalletId          string    `form:"wallet" json:"wallet"`
	Name              string    `form:"name" json:"name"`
	Description       string    `form:"description" json:"description"`
	Value             float64   `form:"value" json:"value"`
	TransactionDate   time.Time `form:"transactionDate" json:"transactionDate"`
}

type UpdateTransactionOutput struct {
	Id    string `form:"id" json:"id"`
	Error error  `form:"error" json:"error"`
}

type DeleteTransactionInput struct {
	Id string `form:"id" json:"id"`
}

type DeleteTransactionOutput struct {
	Id    string `form:"id" json:"id"`
	Error error  `form:"error" json:"error"`
}

type Transaction struct {
	Id                string    `form:"id" json:"id"`
	TransactionTypeId string    `form:"TransactionTypeId" json:"TransactionTypeId"`
	WalletId          string    `form:"wallet" json:"wallet"`
	Name              string    `form:"name" json:"name"`
	Description       string    `form:"description" json:"description"`
	Value             float64   `form:"value" json:"value"`
	TransactionDate   time.Time `form:"transactionDate" json:"transactionDate"`
	CreatedAt         time.Time `form:"createdAt" json:"createdAt"`
}
