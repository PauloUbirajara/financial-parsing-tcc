package models_test

import (
	"testing"
	"time"

	models "financial-parsing/models"
)

func TestTransaction_ShouldBelongToAValidWallet(t *testing.T) {
	transaction, err := models.NewTransaction(nil, "Valid Name", "Valid Description", 123.45, time.Now())

	if err == nil {
		t.Errorf("Did not throw error when creating invalid transaction")
	}

	if transaction != nil {
		t.Errorf("Created invalid transaction with null wallet")
	}
}
