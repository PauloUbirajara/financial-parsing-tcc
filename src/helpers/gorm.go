package helpers

import (
	"fmt"

	configuration "financial-parsing/src/configuration"
	models "financial-parsing/src/domain/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateConnection() (*gorm.DB, error) {
	connectionString := fmt.Sprintf(
		"host=%s user=%s password=%s port=%s sslmode=disable TimeZone=UTC",
		configuration.DatabaseHost,
		configuration.DatabaseUsername,
		configuration.DatabasePassword,
		configuration.DatabasePort,
	)
	db, err := gorm.Open(
		postgres.Open(connectionString),
		&gorm.Config{},
	)

	db.AutoMigrate(&models.Currency{})

	return db, err
}
