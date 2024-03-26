package usecases

type ValidateModel[T interface{}] interface {
	Validate(model T) error
}
