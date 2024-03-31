package validatemodel

import (
	"errors"

	models "financial-parsing/src/domain/models"
	protocols "financial-parsing/src/protocols"
)

type ValidateCurrencyRecord struct {
	UUIDGenerator protocols.UUIDGenerator
}

func (v ValidateCurrencyRecord) Validate(model models.CurrencyRecord) error {
	// ID
	if v.UUIDGenerator.IsValidUUID(model.ID) != nil {
		return errors.New("Invalid ID")
	}

	// Record Date
	if model.RecordDate.IsZero() {
		return errors.New("Invalid Record Date")
	}

	return nil
}
