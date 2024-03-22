package uuidgenerator

import (
	"github.com/google/uuid"
)

type GoogleUUIDGenerator struct{}

func (g GoogleUUIDGenerator) Generate() string {
	uuid := uuid.Must(uuid.NewV7())
	return uuid.String()
}
