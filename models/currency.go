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
