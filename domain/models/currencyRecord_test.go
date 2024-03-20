package models_test

import (
	"testing"
	"time"

	models "financial-parsing/domain/models"
)

func currencyRecordWithMissingCurrency() (*models.CurrencyRecord, error) {
	return models.NewCurrencyRecord(nil, 123.45, time.Now())
}

func validCurrencyRecord() (*models.CurrencyRecord, error) {
	return models.NewCurrencyRecord(&models.Currency{}, 123.45, time.Now())
}

func TestCurrencyRecord_ShouldFailIfCurrencyIsNull(t *testing.T) {
	currencyRecord, err := currencyRecordWithMissingCurrency()

	if err == nil {
		t.Errorf("Did not throw error when creating invalid currencyRecord")
	}

	if currencyRecord != nil {
		t.Errorf("Created invalid currencyRecord with null currency")
	}
}

func TestCurrencyRecord_ShouldContainData(t *testing.T) {
	currencyRecord, _ := validCurrencyRecord()

	if currencyRecord.CreatedAt.IsZero() { t.Errorf("Valid currencyRecord does did not include createdAt when created") }
	if currencyRecord.Currency == nil { t.Errorf("Valid currencyRecord does did not include currency when created") }
	if currencyRecord.RecordDate.IsZero() { t.Errorf("Valid currencyRecord does did not include recordDate when created") }
	if currencyRecord.Value == 0 { t.Errorf("Valid currencyRecord does did not include value when created") }
}

func TestCurrencyRecord_ShouldPassIfValidData(t *testing.T) {
	currencyRecord, err := validCurrencyRecord()

	if err != nil {
		t.Errorf("Threw error when creating valid currency record")
	}

	if currencyRecord == nil {
		t.Errorf("Created null currency record with valid values")
	}
}
