package ms

import (
	"ac/domain/port/in"
	"ac/shared"
	"embed"
	"io/fs"
	"strings"
)

//go:embed assets/*
var files embed.FS

type UsecaseService struct{}

func NewUsecaseService() in.IMSUsecase {
	return &UsecaseService{}
}

func (u *UsecaseService) Create(usecaseName string) error {
	return fs.WalkDir(files, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		fileInfo, content, err := shared.ReadFileWithFileInfo(path, &files)
		if err != nil {
			return err
		}

		projectName := shared.GetProjectName()

		fileInfo.Path = strings.ReplaceAll(fileInfo.Path, "project_name", projectName)
		fileInfo.Path = strings.ReplaceAll(fileInfo.Path, "use_case_name", usecaseName)

		content = strings.ReplaceAll(content, "project_name", projectName)
		content = strings.ReplaceAll(content, "use_case_name", usecaseName)

		content = strings.ReplaceAll(content, "UseCaseName", shared.ToCamelCase(usecaseName))

		err = shared.CreateFile(fileInfo.Path, content)
		if err != nil {
			return err
		}

		return nil
	})
}
