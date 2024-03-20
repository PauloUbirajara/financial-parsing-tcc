package models

import (
	"time"
)

type Wallet struct {
	Id string

	Currency *Currency

	Name        string
	Description string

	CreatedAt time.Time
}
