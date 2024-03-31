package models

import (
	"time"
)

type Wallet struct {
	Id          string    `json:"id"`
	Currency    *Currency `json:"currency"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
}
