package shared

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
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

func RunCommand(command string, args []string, workingDirectory string, printOutput bool) error {
	cmd := exec.Command(command, args...)
	if workingDirectory != "" {
		cmd.Dir = workingDirectory
	}

	output, err := cmd.CombinedOutput()
	if printOutput {
		fmt.Println(string(output))
	}

	return err
}
