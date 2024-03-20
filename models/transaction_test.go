package models_test

import (
	"testing"
	"time"

	models "financial-parsing/models"
)

func transactionWithMissingWallet() (*models.Transaction, error) {
	return models.NewTransaction(nil, "Valid Name", "Valid Description", 123.45, time.Now())
}

func validTransaction() (*models.Transaction, error) {
	return models.NewTransaction(&models.Wallet{}, "Valid Name", "Valid Description", 123.45, time.Now())
}

func TestTransaction_ShouldFailIfWalletIsNull(t *testing.T) {
	transaction, err := transactionWithMissingWallet()

	if err == nil {
		t.Errorf("Did not throw error when creating invalid transaction")
	}

	if transaction != nil {
		t.Errorf("Created invalid transaction with null wallet")
	}
}

func TestTransaction_ShouldPassIfValidData(t *testing.T) {
	transaction, err := validTransaction()

	if err != nil {
		t.Errorf("Threw error when creating valid transaction")
	}

	if transaction == nil {
		t.Errorf("Created null transaction with valid values")
	}
}
