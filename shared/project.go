package shared

import (
	"bufio"
	"os"
	"strings"
)

func GetProjectName() string {
	file, err := os.Open(".go.mod")
	if err != nil {
		return ""
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "module") {
			moduleName := strings.TrimSpace(line[len("module"):])
			return moduleName
		}
	}

	return ""
}
