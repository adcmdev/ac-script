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

func Checkout() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: ac checkout [branch]")
		return
	}

	acGit := shared.NewACGit()
	checkoutService := git.NewCheckoutUsecaseService(acGit)
	err := checkoutService.Checkout(os.Args[2])
	if err != nil {
		fmt.Println(err)
	}
}
