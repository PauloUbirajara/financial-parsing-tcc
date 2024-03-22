package models

import (
	"time"
)

type Currency struct {
	Id             string    `json:"id"`
	Name           string    `json:"name"`
	Representation string    `json:"representation"`
	CreatedAt      time.Time `json:"createdAt"`
}

func NewCurrency(
	name string,
	representation string,
) (*Currency, error) {
	currency := Currency{
		Name:           name,
		Representation: representation,
		CreatedAt:      time.Now(),
	}
	return &currency, nil
}
