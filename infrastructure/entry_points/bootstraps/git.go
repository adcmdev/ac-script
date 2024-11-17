package bootstraps

import (
	"ac/domain/usecase/git"
	"ac/shared"
	"fmt"
	"os"
)

func Commit() {
	if len(os.Args) < 1 {
		fmt.Println("Invalid command")
		return
	}

	acGit := shared.NewACGit()
	commitService := git.NewUsecaseService(acGit)
	err := commitService.Commit(os.Args[2])
	if err != nil {
		fmt.Println(err)
	}
}
