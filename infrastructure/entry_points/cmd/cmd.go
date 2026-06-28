package cmd

import (
	"ac/infrastructure/entry_points/bootstraps"
	"fmt"
	"os"
)

func StartCliApplication() {
	if len(os.Args) < 2 {
		fmt.Println(`
 _______  _______
 $$$$$$\   $$$$$$\  
$$  __$$\ $$  __$$\ 
$$ /  $$ |$$ /  \__|
$$$$$$$$ |$$ |      
$$  __$$ |$$ |      
$$ |  $$ |$$ |  $$\ 
$$ |  $$ |\$$$$$$  |
\__|  \__| \______/
		`)
		fmt.Println("Usage: ac [command]")
		fmt.Println("Commands:")
		fmt.Println("  usecase   Create a new usecase")
		fmt.Println("  mobile    Deploy the mobile app")
		fmt.Println("  new       Create a new project")
		fmt.Println("  commit    Commit and push the changes")
		fmt.Println("  checkout  Checkout a branch, creating it if it does not exist")

		return
	}

	commandType := os.Args[1]

	switch commandType {
	case "mobile":
		bootstraps.DeployLocal()
	case "commit":
		bootstraps.Commit()
	case "checkout":
		bootstraps.Checkout()
	case "new":
		bootstraps.NewProject()
	default:
		fmt.Println("Invalid command type")
	}
}
