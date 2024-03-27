package databaseadapter_test

import (
	"log"
	"testing"

	databaseadapter "financial-parsing/src/utils/databaseAdapter"

	"database/sql"
	"database/sql/driver"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewMockDB() (*sql.DB, *gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()

	if err != nil {
		log.Panic("Failed when creating a stub database connection", err)
	}

	gormDB, err := gorm.Open(
		postgres.New(postgres.Config{
			Conn:       db,
			DriverName: "postgres",
		}),
		&gorm.Config{},
	)
	if err != nil {
		log.Panic("Failed when opening gorm database", err)
	}

	return db, gormDB, mock
}

type TestModel struct {
	ID   string
	Name string
}

type GormDatabaseAdapterTestData struct {
	sut  databaseadapter.GormDatabaseAdapter[TestModel]
	mock sqlmock.Sqlmock
	db   *sql.DB
}

func NewGormDatabaseAdapterTestData() *GormDatabaseAdapterTestData {
	db, gormDb, mock := NewMockDB()

	return &GormDatabaseAdapterTestData{
		sut: databaseadapter.GormDatabaseAdapter[TestModel]{
			Connection: gormDb,
		},
		mock: mock,
		db:   db,
	}
}

func TestGormDatabaseAdapterShouldPassOnGetAll(t *testing.T) {
	testData := NewGormDatabaseAdapterTestData()
	defer testData.db.Close()

	rows := testData.mock.
		NewRows([]string{"id", "name"}).
		AddRows(
			[]driver.Value{"valid id", "valid name"},
			[]driver.Value{"valid id 2", "valid name 2"},
		)

	testData.mock.
		ExpectQuery(`SELECT (.+) FROM "test_models"`).
		WillReturnRows(rows)

	models, err := testData.sut.GetAll()

	if err != nil {
		t.Fatal("Error when getting all test models", err)
	}

	if len(*models) != 2 {
		t.Fatal("Did not get all test models", models)
	}
}

func TestGormDatabaseAdapterShouldPassOnGetById(t *testing.T) {
	testData := NewGormDatabaseAdapterTestData()
	defer testData.db.Close()

	rows := testData.mock.
		NewRows([]string{"id", "name"}).
		AddRows(
			[]driver.Value{"valid id", "valid name"},
			[]driver.Value{"valid id 2", "valid name 2"},
		)

	testData.mock.
		ExpectQuery(`SELECT (.+) FROM "test_models" WHERE id = ?`).
		WillReturnRows(rows)

	model, err := testData.sut.GetById("valid id")

	if err != nil {
		t.Fatal("Error when getting test model by id", err)
	}

	if model.ID != "valid id" {
		t.Fatal("Did not get test model by id", model, err, model.ID)
	}
}

func TestGormDatabaseAdapterShouldPassOnCreate(t *testing.T) {
	testData := NewGormDatabaseAdapterTestData()
	defer testData.db.Close()

	testData.mock.ExpectBegin()
	testData.mock.
		ExpectExec(`INSERT`).
		WithArgs("valid id 2", "valid name 2").
		WillReturnResult(sqlmock.NewResult(0, 1))
	testData.mock.ExpectCommit()

	testModel := TestModel{
		ID:   "valid id 2",
		Name: "valid name 2",
	}
	model, err := testData.sut.Create(&testModel, []string{"ID", "Name"})

	if err != nil {
		t.Fatal("Error when creating test model", err, model)
	}

	if model.ID != "valid id 2" {
		t.Fatal("Did not create test model", model)
	}
}

func TestGormDatabaseAdapterShouldPassOnDeleteWithSingleID(t *testing.T) {
	testData := NewGormDatabaseAdapterTestData()
	defer testData.db.Close()

	testData.mock.ExpectBegin()
	testData.mock.
		ExpectExec(`DELETE`).
		WithArgs("valid id 2").
		WillReturnResult(sqlmock.NewResult(0, 1))
	testData.mock.ExpectCommit()

	err := testData.sut.DeleteByIds([]string{"valid id 2"})

	if err != nil {
		t.Fatal("Error when deleting test model with single id", err)
	}
}

func TestGormDatabaseAdapterShouldPassOnDeleteWithMultipleIDS(t *testing.T) {
	testData := NewGormDatabaseAdapterTestData()
	defer testData.db.Close()

	testData.mock.ExpectBegin()
	testData.mock.
		ExpectExec(`DELETE`).
		WithArgs("valid id", "valid id 2").
		WillReturnResult(sqlmock.NewResult(0, 2))
	testData.mock.ExpectCommit()

	err := testData.sut.DeleteByIds([]string{"valid id", "valid id 2"})

	if err != nil {
		t.Fatal("Error when deleting test model with multiple ids", err)
	}
}

func TestGormDatabaseAdapterShouldPassOnUpdate(t *testing.T) {
	testData := NewGormDatabaseAdapterTestData()
	defer testData.db.Close()

	// Expected content to be updated
	rows := testData.mock.
		NewRows([]string{"id", "name"}).
		AddRows(
			[]driver.Value{"valid id", "valid name"},
		)

	// Should select and update
	testData.mock.
		ExpectQuery(`SELECT (.+) FROM "test_models" WHERE id = ?`).
		WithArgs("valid id", 1).
		WillReturnRows(rows)
	testData.mock.ExpectBegin()
	testData.mock.
		ExpectExec(`UPDATE`).
		WithArgs("valid name 2", "valid id").
		WillReturnResult(sqlmock.NewResult(0, 1))
	testData.mock.ExpectCommit()
	updatedTestModel := TestModel{
		ID:   "valid id 2",
		Name: "valid name 2",
	}
	fields := []string{"Name"}
	model, err := testData.sut.UpdateById("valid id", &updatedTestModel, fields)

	if err != nil {
		t.Fatal("Error when updating test model", err, model)
	}

	if model.Name != "valid name 2" {
		t.Fatal("Did not update test model", model)
	}
}
