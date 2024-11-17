package bootstraps

import (
	"ac/domain/usecase/ms"
	"fmt"
	"os"
)

func CreateUsecase() {
	if len(os.Args) < 1 {
		fmt.Println("Invalid command")
		return
	}

	msService := ms.NewUsecaseService()

	useCaseName := os.Args[2]

	err := msService.Create(useCaseName)
	if err != nil {
		fmt.Println(err)
	}
}
