//go:build exclude

/*
{
	"package": "in",
	"path": "domain/ports/in/use_case_name_usecase.go"
}
*/

package in

type IUseCaseNameUseCase interface {
	Execute() (interface{}, error)
}
