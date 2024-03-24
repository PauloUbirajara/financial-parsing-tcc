package modelcommand

import (
	protocols "financial-parsing/src/protocols"
)

type GetAllModelsCommand[Input, Output interface{}] struct {
	DatabaseAdapter protocols.DatabaseAdapter[Input, Output]
}

func (g GetAllModelsCommand[Input, Output]) Run(input Input) (*Output, error) {
	transactions, err := g.DatabaseAdapter.GetAll(input)
	return transactions, err
}
