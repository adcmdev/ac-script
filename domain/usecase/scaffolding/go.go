package scaffolding

import (
	"ac/shared"
	"fmt"
	"os"
	"strings"
)

func (u *UsecaseService) startGOProject(projectName string) (err error) {
	err = u.git.Clone("https://github.com/adcmdev/go-template.git", "")
	if err != nil {
		return
	}

	err = shared.Rename("go-template", projectName)
	if err != nil {
		return
	}

	err = os.RemoveAll(projectName + "/.git")

	dbPort, err := shared.AskForInput("Enter the port of the database: ")
	if err != nil {
		return
	}

	pInternalPort, err := shared.AskForInput("Enter the internal port of the project: ")
	if err != nil {
		return
	}

	pExternalPort, err := shared.AskForInput("Enter the external port of the project: ")
	if err != nil {
		return
	}

	replaceMap := map[string]string{
		"replace_this_with_project_name":                 projectName,
		"REPLACE_THIS_WITH_PROJECT_NAME":                 strings.ToUpper(projectName),
		"REPLACE_THIS_WITH_PROJECT_DB_PORT":              dbPort,
		"REPLACE_THIS_WITH_PROJECT_TOKEN_SECRET":         shared.RandomString(32),
		"REPLACE_THIS_WITH_PROJECT_REFRESH_TOKEN_SECRET": shared.RandomString(32),
		"REPLACE_THIS_WITH_PROJECT_INTERNAL_PORT":        pInternalPort,
		"REPLACE_THIS_WITH_PROJECT_EXTERNAL_PORT":        pExternalPort,
	}

	err = u.replaceTextInFiles(projectName, replaceMap)
	if err != nil {
		return
	}

	err = u.git.Init(projectName, fmt.Sprintf("./%s", projectName))
	if err != nil {
		return
	}

	return
}
