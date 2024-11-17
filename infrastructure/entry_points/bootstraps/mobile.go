package bootstraps

import (
	"ac/domain/entities"
	"ac/domain/usecase/mobile"
	"fmt"
	"os"
)

func DeployLocal() {
	if len(os.Args) < 1 {
		fmt.Println("Invalid command")
		return
	}

	mobileService := mobile.NewUsecaseService()

	arg := ""

	if len(os.Args) >= 5 {
		arg = os.Args[4]
	} else {
		arg = "bugfix"
	}

	useCaseName := entities.TulReleaseTypeFromString(arg)

	err := mobileService.DeployLocal(useCaseName)
	if err != nil {
		fmt.Println(err)
	}
}
