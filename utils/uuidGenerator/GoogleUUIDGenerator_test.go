package uuidgenerator_test

import (
	"testing"

	protocols "financial-parsing/protocols"
	uuidGenerator "financial-parsing/utils/uuidGenerator"
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

func TestGoogleUUIDGenerator_ShouldGenerateUUID(t *testing.T) {
	testData := NewUUIDGeneratorTestData()
	uuid := testData.sut.Generate()

	if uuid == "" {
		t.Errorf("Did not generate an UUID")
	}
}
