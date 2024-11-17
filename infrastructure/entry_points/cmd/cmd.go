package cmd

import (
	"ac/infrastructure/entry_points/bootstraps"
	"fmt"
	"os"
)

func StartCliApplication() {
	if len(os.Args) < 2 {
		fmt.Println(`
████████╗██╗   ██╗██╗     ██╗ █████╗
╚══██╔══╝██║   ██║██║     ██║██╔══██╗
   ██║   ██║   ██║██║     ██║███████║
   ██║   ██║   ██║██║     ██║██╔══██║
   ██║   ╚██████╔╝███████╗██║██║  ██║
   ╚═╝    ╚═════╝ ╚══════╝╚═╝╚═╝  ╚═╝

		`)
		fmt.Println("Usage: ac [command]")
		fmt.Println("Commands:")
		fmt.Println("  plugin   Create a new plugin")
		fmt.Println("  usecase  Create a new usecase")
		fmt.Println("  mobile   Deploy the mobile app")

		return
	}

	commandType := os.Args[1]

	switch commandType {
	case "usecase":
		bootstraps.CreateUsecase()
	case "mobile":
		bootstraps.DeployLocal()
	case "commit":
		bootstraps.Commit()
	default:
		fmt.Println("Invalid command type")
	}
}
