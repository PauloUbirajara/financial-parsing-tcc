package protocols

type DatabaseAdapter[T interface{}] interface {
	GetAll() (*[]T, error)
	GetById(id string) (*T, error)
	Create(model *T, fieldNames []string) (*T, error)
	DeleteByIds(ids []string) error
	UpdateById(id string, updated *T) (*T, error)
}
