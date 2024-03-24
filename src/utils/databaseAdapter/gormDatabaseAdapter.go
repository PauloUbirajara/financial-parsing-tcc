package databaseadapter

import (
	"gorm.io/gorm"
)

type GormDatabaseAdapter[T interface{}] struct {
	Connection *gorm.DB
}

func (g GormDatabaseAdapter[T]) GetAll() (*[]T, error) {
	var models []T
	result := g.Connection.Find(&models)
	return &models, result.Error
}

// func (p PostgresDatabaseAdapter) Create(id string, model any) error {}
// func (p PostgresDatabaseAdapter) DeleteById(id string) (*any, error) {}
// func (p PostgresDatabaseAdapter) GetById(id string) (*any, error)                 {}
// func (p PostgresDatabaseAdapter) UpdateById(id string, updated any) (*any, error) {}
