package scaffolding

import (
	"ac/domain/port/in"
	"ac/shared"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type ProjectTypes string

func (p ProjectTypes) String() string {
	return string(p)
}

const (
	Flutter ProjectTypes = "Flutter"
	GO      ProjectTypes = "GO"
	React   ProjectTypes = "React"
)

type UsecaseService struct {
	git in.IGit
}

func NewUsecaseService(git in.IGit) in.IScaffoldingUsecase {
	return &UsecaseService{
		git: git,
	}
}

func (u *UsecaseService) StartNewProject() error {
	projectName, err := shared.AskForInput("Enter the name of the project: ")
	if err != nil {
		return err
	}

	result, err := shared.ShowMultipleChoiceSelector(
		"Select the type of project you want to create",
		[]shared.MultipleChoiceItem{
			{Text: Flutter.String(), Color: shared.Blue},
			{Text: GO.String(), Color: shared.DarkBlue},
			{Text: React.String(), Color: shared.Cyan},
		},
	)

	if err != nil {
		return err
	}

	fmt.Println(fmt.Sprintf("Selected %s", result))

	projectType := ProjectTypes(result)
	err = u.StartNewProjectWithType(projectType, projectName)
	if err != nil {
		return err
	}

	return nil
}

func (u *UsecaseService) StartNewProjectWithType(projectType ProjectTypes, projectName string) (err error) {
	fmt.Println(fmt.Sprintf("Creating new %s project", projectType.String()))

	switch projectType {
	case Flutter:
		err = u.startFlutterProject()
		break
	case GO:
		err = u.startGOProject(projectName)
		break
	case React:
		err = u.startReactProject()
		break
	}

	if err != nil {
		fmt.Println(fmt.Sprintf("Error creating new %s project. ERROR: %s", projectType, err.Error()))
		return err
	}

	fmt.Println("Project created successfully")

	return nil
}

func (u *UsecaseService) replaceTextInFiles(dir string, replacements map[string]string) error {
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			data, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			updatedData := string(data)
			for oldText, newText := range replacements {
				updatedData = strings.ReplaceAll(updatedData, oldText, newText)
			}

			if updatedData != string(data) {
				err = os.WriteFile(path, []byte(updatedData), info.Mode())
				if err != nil {
					return err
				}
			}
		}
		return nil
	})
	return err
}
