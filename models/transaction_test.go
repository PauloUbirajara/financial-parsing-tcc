package models_test

import (
	"testing"
	"time"

	models "financial-parsing/models"
)

func TestTransaction_ShouldFailIfWalletIsNull(t *testing.T) {
	transaction, err := models.NewTransaction(nil, "Valid Name", "Valid Description", 123.45, time.Now())

	if err == nil {
		t.Errorf("Did not throw error when creating invalid transaction")
	}

	if transaction != nil {
		t.Errorf("Created invalid transaction with null wallet")
	}
}

func TestTransaction_ShouldPassIfValidData(t *testing.T) {
	transaction, err := models.NewTransaction(&models.Wallet{}, "Valid Name", "Valid Description", 123.45, time.Now())

	if err != nil {
		t.Errorf("Threw error when creating valid transaction")
	}

	if transaction == nil {
		t.Errorf("Created null transaction with valid values")
	}
}
