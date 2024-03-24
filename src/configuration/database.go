package configuration

import "os"

var (
	DatabaseHost     = os.Getenv("POSTGRES_HOST")
	DatabaseUsername = os.Getenv("POSTGRES_USER")
	DatabasePassword = os.Getenv("POSTGRES_PASSWORD")
	DatabasePort     = os.Getenv("POSTGRES_PORT")
)
