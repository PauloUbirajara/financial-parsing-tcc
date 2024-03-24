package helpers

import (
	"financial-parsing/src/domain/models"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateConnection() (*gorm.DB, error) {
	var (
		host     = os.Getenv("POSTGRES_HOST")
		username = os.Getenv("POSTGRES_USER")
		password = os.Getenv("POSTGRES_PASSWORD")
		database = os.Getenv("POSTGRES_DATABASE")
		port     = os.Getenv("POSTGRES_PORT")
	)

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s port=%s sslmode=disable TimeZone=UTC",
		host,
		username,
		password,
		port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&models.Currency{})

	return db, err
}
