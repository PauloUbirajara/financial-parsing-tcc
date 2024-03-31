package validatemodel

import (
	"errors"
	"regexp"

	models "financial-parsing/src/domain/models"
	protocols "financial-parsing/src/protocols"
)

type ValidateCurrency struct {
	UUIDGenerator protocols.UUIDGenerator
}

func (v ValidateCurrency) Validate(model models.Currency) error {
	// ID
	if v.UUIDGenerator.IsValidUUID(model.ID) != nil {
		return errors.New("Invalid ID")
	}

	// Name
	nameRegex, _ := regexp.Compile(`^\S[\S, ]{1,62}\S$`)
	if !nameRegex.MatchString(model.Name) {
		return errors.New("Invalid Name")
	}

	// Representation
	representationRegex, _ := regexp.Compile("^[A-Z]{1,5}$")
	if !representationRegex.MatchString(model.Representation) {
		return errors.New("Invalid Representation")
	}

	return nil
}
