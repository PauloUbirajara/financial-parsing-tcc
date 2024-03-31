package validatemodel_test

import (
	"errors"
	"testing"
	"time"

	models "financial-parsing/src/domain/models"
	usecases "financial-parsing/src/domain/usecases"

	validatemodel "financial-parsing/src/data/usecases/validateModel"

	"gorm.io/gorm"
)

type ValidateCurrencyTestData struct {
	sut usecases.ValidateModel[models.Currency]
}

type UUIDGeneratorStub struct{}

func (u UUIDGeneratorStub) Generate() string {
	return ""
}

func (u UUIDGeneratorStub) IsValidUUID(id string) error {
	if id != "valid id" {
		return errors.New("invalid id")
	}
	return nil
}

func newValidCurrency() models.Currency {
	return models.Currency{
		ID:             "valid id",
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		DeletedAt:      gorm.DeletedAt{},
		Name:           "valid name",
		Representation: "REPR",
	}
}

func validateCurrencyTestData() ValidateCurrencyTestData {
	return ValidateCurrencyTestData{
		sut: validatemodel.ValidateCurrency{
			UUIDGenerator: UUIDGeneratorStub{},
		},
	}
}

func TestValidateCurrencyShouldFailOnEmptyName(t *testing.T) {
	currency := newValidCurrency()
	currency.Name = ""

	testData := validateCurrencyTestData()

	if testData.sut.Validate(currency) == nil {
		t.Errorf("Invalid Currency considered valid with empty name")
	}
}

func TestValidateCurrencyShouldFailOnInvalidName(t *testing.T) {
	currency := newValidCurrency()
	currency.Name = " invalid name"

	testData := validateCurrencyTestData()

	if testData.sut.Validate(currency) == nil {
		t.Errorf("Invalid Currency considered valid with invalid name")
	}
}

func TestValidateCurrencyShouldFailOnEmptyRepresentation(t *testing.T) {
	currency := newValidCurrency()
	currency.Representation = ""

	testData := validateCurrencyTestData()

	if testData.sut.Validate(currency) == nil {
		t.Errorf("Invalid Currency considered valid with empty representation")
	}
}

func TestValidateCurrencyShouldFailOnInvalidRepresentation(t *testing.T) {
	currency := newValidCurrency()
	currency.Representation = "abc"

	testData := validateCurrencyTestData()

	if testData.sut.Validate(currency) == nil {
		t.Errorf("Invalid Currency considered valid with invalid representation")
	}
}

func TestValidateCurrencyShouldFailOnInvalidID(t *testing.T) {
	currency := newValidCurrency()
	currency.ID = "invalid id"

	testData := validateCurrencyTestData()

	if testData.sut.Validate(currency) == nil {
		t.Errorf("Invalid Currency considered valid with empty id")
	}
}

func TestValidateCurrencyShouldPassOnValidCurrency(t *testing.T) {
	currency := newValidCurrency()

	testData := validateCurrencyTestData()

	if testData.sut.Validate(currency) != nil {
		t.Errorf("Valid Currency considered invalid")
	}
}
