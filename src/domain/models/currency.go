package models

import (
	"time"
)

type Currency struct {
	Id string

	Name           string
	Representation string

	CreatedAt time.Time
}

func NewCurrency(
	name string,
	representation string,
) (*Currency, error) {
	currency := Currency{
		Name: name,
		Representation: representation,

		CreatedAt: time.Now(),
	}
	return &currency, nil
}
