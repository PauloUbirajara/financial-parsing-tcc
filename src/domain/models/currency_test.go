package models_test

import (
	"testing"

	models "financial-parsing/src/domain/models"
)

func validCurrency() (*models.Currency, error) {
	return models.NewCurrency("Valid Name", "VALIDREPRESENTATION")
}

func TestCurrency_ShouldContainData(t *testing.T) {
	currency, _ := validCurrency()

	if currency.CreatedAt.IsZero() { t.Errorf("Valid currency does did not include createdAt when created") }
	if currency.Name == "" { t.Errorf("Valid currency does did not include name when created") }
	if currency.Representation == "" { t.Errorf("Valid currency does did not include representation when created") }
}

func TestCurrency_ShouldPassIfValidData(t *testing.T) {
	currency, err := validCurrency()

	if err != nil {
		t.Errorf("Threw error when creating valid currency record")
	}

	if currency == nil {
		t.Errorf("Created null currency record with valid values")
	}
}
