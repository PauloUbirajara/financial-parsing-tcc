package usecases

type ModelCommand[Input, Output interface{}] interface {
	Run(input Input) (*Output, error)
}
