//go:build exclude

/*
{
	"package": "use_case_name",
	"path": "infrastructure/entry_points/rest_api/use_case_name.go"
}
*/

package rest_api

import (
	"net/http"
	"project_name/domain/ports/in"

	"github.com/labstack/echo"
)

type UseCaseNameHandler struct {
	useCase in.IUseCaseNameUseCase
}

func NewUseCaseNameHandler(useCase in.IUseCaseNameUseCase) *UseCaseNameHandler {
	return &UseCaseNameHandler{
		useCase: useCase,
	}
}

func (m *UseCaseNameHandler) Execute(c echo.Context) error {
	response, err := m.useCase.Execute()

	if err != nil {
		return c.JSON(
			http.StatusUnprocessableEntity,
			map[string]string{
				"error": err.Error(),
			},
		)
	}

	return c.JSON(http.StatusOK, response)
}
