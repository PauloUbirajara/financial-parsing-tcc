package uuidgenerator_test

import (
	"testing"

	protocols "financial-parsing/src/protocols"

	uuidGenerator "financial-parsing/src/utils/uuidGenerator"
)

type UUIDGeneratorTestData struct {
	sut protocols.UUIDGenerator
}

func NewUUIDGeneratorTestData() *UUIDGeneratorTestData {
	testData := UUIDGeneratorTestData{
		sut: &uuidGenerator.GoogleUUIDGenerator{},
	}

	return &testData
}

func TestGoogleUUIDGeneratorShouldGenerateUUID(t *testing.T) {
	testData := NewUUIDGeneratorTestData()
	uuid := testData.sut.Generate()

	if uuid == "" {
		t.Errorf("Did not generate an UUID")
	}
}

func TestGoogleUUIDGeneratorShouldFailWhenValidatingInvalidUUID(t *testing.T) {
	testData := NewUUIDGeneratorTestData()
	invalidUUID := "invalid"
	err := testData.sut.IsValidUUID(invalidUUID)

	if err == nil {
		t.Errorf("SUT did not throw error when validating invalid UUID")
	}
}

func TestGoogleUUIDGeneratorShouldPassWhenValidatingValidUUID(t *testing.T) {
	testData := NewUUIDGeneratorTestData()
	validUUID := "11111111-1111-1111-1111-111111111111"
	err := testData.sut.IsValidUUID(validUUID)

	if err != nil {
		t.Errorf("SUT threw error when validating valid UUID")
	}
}
