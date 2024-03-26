package validatemodel

import (
	"regexp"

	models "financial-parsing/src/domain/models"
	"financial-parsing/src/protocols"
)

type ValidateCurrency struct {
	Currency      models.Currency
	UUIDGenerator protocols.UUIDGenerator
}

func (v ValidateCurrency) Validate() bool {
	// ID
	if v.UUIDGenerator.IsValidUUID(v.Currency.ID) != nil {
		return false
	}

	// Name
	nameRegex, _ := regexp.Compile(".{1,64}")
	if !nameRegex.MatchString(v.Currency.Name) {
		return false
	}

	// Representation
	representationRegex, _ := regexp.Compile(".{0,256}")
	if !representationRegex.MatchString(v.Currency.Representation) {
		return false
	}

	return true
}
