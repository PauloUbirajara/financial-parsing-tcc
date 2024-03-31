package uuidgenerator

import (
	"github.com/google/uuid"
)

type GoogleUUIDGenerator struct{}

func (g GoogleUUIDGenerator) Generate() string {
	uuid := uuid.Must(uuid.NewV7())
	return uuid.String()
}

func (g GoogleUUIDGenerator) IsValidUUID(id string) error  {
	return uuid.Validate(id)
}
