//go:build exclude

/*
{
	"package": "usecase",
	"path": "domain/usecase/use_case_name.go"
}
*/

package usecase

import (
	"project_name/domain/ports/in"
)

type use_case_name struct{}

func NewMarketing() in.IUseCaseNameUseCase {
	return &use_case_name{}
}

func (m *use_case_name) Execute() (interface{}, error) {
	return nil, nil
}
