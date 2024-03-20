package models_test

import (
	"testing"

	models "financial-parsing/models"
)

func walletWithMissingCurrency() (*models.Wallet, error) {
	return models.NewWallet(nil, "Valid Name", "Valid Description")
}

func validWallet() (*models.Wallet, error) {
	return models.NewWallet(&models.Currency{}, "Valid Name", "Valid Description")
}

func TestWallet_ShouldFailIfCurrencyIsNull(t *testing.T) {
	wallet, err := walletWithMissingCurrency()

	if err == nil {
		t.Errorf("Did not throw error when creating invalid wallet")
	}

	if wallet != nil {
		t.Errorf("Created invalid wallet with null currency")
	}
}

func TestWallet_ShouldPassIfValidData(t *testing.T) {
	wallet, err := validWallet()

	if err != nil {
		t.Errorf("Threw error when creating valid wallet")
	}

	if wallet == nil {
		t.Errorf("Created null wallet with valid values")
	}
}
