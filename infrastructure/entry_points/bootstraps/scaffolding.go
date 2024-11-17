package bootstraps

import (
	"ac/domain/usecase/scaffolding"
	"ac/shared"
	"fmt"
)

func NewProject() {
	acGit := shared.NewACGit()
	scaffoldingService := scaffolding.NewUsecaseService(acGit)

	err := scaffoldingService.StartNewProject()
	if err != nil {
		fmt.Println(err)
	}
}
