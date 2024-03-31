package helpers

import (
	"fmt"

	models "financial-parsing/src/domain/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateConnection(host, username, password, port string) (*gorm.DB, error) {
	connectionString := fmt.Sprintf(
		"host=%s user=%s password=%s port=%s sslmode=disable TimeZone=UTC",
		host,
		username,
		password,
		port,
	)
	db, err := gorm.Open(
		postgres.Open(connectionString),
		&gorm.Config{},
	)

	db.AutoMigrate(
		&models.User{},
		&models.Currency{},
		&models.Currency_User{},
		&models.CurrencyRecord{},
		&models.Currency_CurrencyRecord{},
	)

	return db, err
}
