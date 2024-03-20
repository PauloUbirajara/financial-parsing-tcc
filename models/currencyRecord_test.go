package models_test

import (
	"testing"
	"time"

	models "financial-parsing/models"
)

func currencyRecordWithMissingCurrency() (*models.CurrencyRecord, error) {
	return models.NewCurrencyRecord(nil, "Valid Name", 123.45, time.Now())
}

func validCurrencyRecord() (*models.CurrencyRecord, error) {
	return models.NewCurrencyRecord(&models.Currency{}, "Valid Name", 123.45, time.Now())
}

func TestCurrencyRecord_ShouldFailIfCurrencyIsNull(t *testing.T) {
	currencyRecord, err := models.NewCurrencyRecord(nil, "Valid Name", 123.45, time.Now())

	if err == nil {
		t.Errorf("Did not throw error when creating invalid currencyRecord")
	}

	if currencyRecord != nil {
		t.Errorf("Created invalid currencyRecord with null currency")
	}
}

func TestCurrencyRecord_ShouldPassIfValidData(t *testing.T) {
	currencyRecord, err := models.NewCurrencyRecord(&models.Currency{}, "Valid Name", 123.45, time.Now())

	if err != nil {
		t.Errorf("Threw error when creating valid currency record")
	}

	if currencyRecord == nil {
		t.Errorf("Created null currency record with valid values")
	}
}
