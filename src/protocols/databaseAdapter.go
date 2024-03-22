package protocols

type DatabaseAdapter interface {
	Create(id string, model any) error
	DeleteById(id string) (*any, error)
	GetAll() (*map[string]any, error)
	GetById(id string) (*any, error)
	UpdateById(id string, updated any) (*any, error)
}
