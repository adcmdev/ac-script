package shared

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreateFolder(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

func CreateFile(path, content string) error {
	dir := filepath.Dir(path)
	if err := CreateFolder(dir); err != nil {
		return fmt.Errorf("error al crear las carpetas: %w", err)
	}

	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("error al crear el archivo: %w", err)
	}

	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return fmt.Errorf("error al escribir en el archivo: %w", err)
	}

	return nil
}

func Rename(source, destination string) error {
	return os.Rename(source, destination)
}
