package databaseadapter

import (
	"errors"
)

type LocalDatabaseAdapter struct {
	Values map[string]any
}

func (l LocalDatabaseAdapter) Create(id string, model any) error {
	_, exists := l.Values[id]

	if exists {
		return errors.New("Error when creating, model already exists")
	}
	l.Values[id] = model
	return nil
}

func (l LocalDatabaseAdapter) DeleteById(id string) (*any, error) {
	deleted, exists := l.Values[id]

	if !exists {
		return nil, nil
	}

	delete(l.Values, id)
	return &deleted, nil
}

func (l LocalDatabaseAdapter) GetAll() (*map[string]any, error) {
	return &l.Values, nil
}

func (l LocalDatabaseAdapter) GetById(id string) (*any, error) {
	model, exists := l.Values[id]

	if !exists {
		return nil, errors.New("Error when getting by id, model not found")
	}

	return &model, nil
}

func (l LocalDatabaseAdapter) UpdateById(id string, updated any) (*any, error) {
	_, exists := l.Values[id]

	if !exists {
		return nil, errors.New("Error when updating, model not found")
	}

	l.Values[id] = updated
	return &updated, nil
}
