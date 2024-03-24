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

func (g GormDatabaseAdapter[T]) GetById(id string) (*T, error) {
	var model T
	result := g.Connection.First(&model, "id = ?", id)
	return &model, result.Error
}

// func (g GormDatabaseAdapter) Create(id string, model any) error {}
// func (g GormDatabaseAdapter) DeleteById(id string) (*any, error) {}
// func (g GormDatabaseAdapter) UpdateById(id string, updated any) (*any, error) {}
