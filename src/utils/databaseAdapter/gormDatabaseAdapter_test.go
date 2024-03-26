package databaseadapter_test

import (
	"database/sql/driver"
	databaseadapter "financial-parsing/src/utils/databaseAdapter"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewMockDB() (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Panic("Failed when creating a stub database connection", err)
	}

	gormDB, err := gorm.Open(
		postgres.New(postgres.Config{
			Conn: db,
		}),
		&gorm.Config{},
	)
	if err != nil {
		log.Panic("Failed when opening gorm database", err)
	}

	return gormDB, mock
}

type TestModel struct {
	ID   string
	Name string
}

type GormDatabaseAdapterTestData struct {
	sut  databaseadapter.GormDatabaseAdapter[TestModel]
	mock sqlmock.Sqlmock
}

func NewGormDatabaseAdapterTestData() *GormDatabaseAdapterTestData {
	gormDb, mock := NewMockDB()

	return &GormDatabaseAdapterTestData{
		sut: databaseadapter.GormDatabaseAdapter[TestModel]{
			Connection: gormDb,
		},
		mock: mock,
	}
}

func TestGormDatabaseAdapterShouldPassOnGetAll(t *testing.T) {
	testData := NewGormDatabaseAdapterTestData()

	valuesToAdd := [][]driver.Value{
		{"valid id", "valid name"},
		{"valid id 2", "valid name 2"},
	}
	rows := testData.mock.NewRows([]string{"id", "name"}).AddRows(valuesToAdd[0], valuesToAdd[1])
	testData.mock.ExpectQuery(`^SELECT (.+) FROM "test_models"$`).WillReturnRows(rows)

	models, err := testData.sut.GetAll()
	if err != nil {
		t.Fatal("Error in finding test model", err)
	}

	if len(*models) != 2 {
		t.Fatal("Failed to return same test model", models)
	}
}
