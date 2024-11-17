package shared

import (
	"ac/domain/port/in"
	"fmt"
	"os/exec"
	"strings"
)

type ACGit struct {
}

func NewACGit() in.IGit {
	return &ACGit{}
}

func (t *ACGit) RunCommand(command string, args []string, workingDirectory string, printOutput bool) bool {
	cmd := exec.Command(command, args...)
	if workingDirectory != "" {
		cmd.Dir = workingDirectory
	}

	output, err := cmd.CombinedOutput()
	if printOutput {
		fmt.Println(string(output))
	}

	return err == nil
}

func (t *ACGit) Add(files, workingDirectory string, filesToExclude []string) {
	if files == "" {
		files = "."
	}

	t.RunCommand("git", []string{"add", files}, workingDirectory, true)

	if len(filesToExclude) > 0 {
		for _, file := range filesToExclude {
			t.RunCommand("git", []string{"reset", file}, workingDirectory, true)
		}
	}
}

func (t *ACGit) Commit(commit, workingDirectory string) {
	t.RunCommand("git", []string{"commit", "-m", commit}, workingDirectory, true)
}

func (t *ACGit) Push(branch, workingDirectory string, force bool, newBranch bool) bool {
	if branch == "" {
		branch = t.CurrentBranch(workingDirectory)
	}

	branchExists := t.BranchExists(branch, workingDirectory)

	if !branchExists {
		newBranch = true
	}

	args := []string{"push"}
	if newBranch {
		args = append(args, "--set-upstream")
	}

	args = append(args, "origin", branch)
	if force {
		args = append(args, "--force")
	}

	return t.RunCommand("git", args, workingDirectory, true)
}

func (t *ACGit) CurrentBranch(workingDirectory string) string {
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	if workingDirectory != "" {
		cmd.Dir = workingDirectory
	}

	output, err := cmd.Output()
	if err != nil {
		return ""
	}

	return strings.TrimSpace(string(output))
}

func (t *ACGit) BranchExists(branch, workingDirectory string) bool {
	cmd := exec.Command("git", "rev-parse", "--verify", branch)
	if workingDirectory != "" {
		cmd.Dir = workingDirectory
	}

	err := cmd.Run()
	return err == nil
}
