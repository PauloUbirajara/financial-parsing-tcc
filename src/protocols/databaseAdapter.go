package protocols

type DatabaseAdapter[T interface{}] interface {
	GetAll() (*[]T, error)
	GetById(id string) (*T, error)
	Create(model *T, fieldNames []string) (*T, error)
	// DeleteById(id string) (*any, error)
	// GetById(id string) (*any, error)
	// UpdateById(id string, updated any) (*any, error)
}
