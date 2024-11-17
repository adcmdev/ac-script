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

func (t *ACGit) Init(repoName, workingDirectory string) (err error) {
	err = RunCommand("gh", []string{"repo", "create", "--private", repoName}, workingDirectory, true)
	if err != nil {
		return
	}

	err = RunCommand("git", []string{"init"}, workingDirectory, true)
	if err != nil {
		return
	}

	t.Add(".", workingDirectory, []string{})
	t.Commit("Initial commit", workingDirectory)
	err = RunCommand("git", []string{"branch", "-M", "main"}, workingDirectory, true)
	if err != nil {
		return
	}

	err = RunCommand(
		"git",
		[]string{"remote", "add", "origin", fmt.Sprintf("https://github.com/adcmdev/%s.git", repoName)},
		workingDirectory,
		true,
	)

	if err != nil {
		return
	}

	err = RunCommand("git", []string{"push", "--set-upstream", "origin", "main"}, workingDirectory, true)

	return
}

func (t *ACGit) Add(files, workingDirectory string, filesToExclude []string) {
	if files == "" {
		files = "."
	}

	err := RunCommand("git", []string{"add", files}, workingDirectory, true)
	if err != nil {
		return
	}

	if len(filesToExclude) > 0 {
		for _, file := range filesToExclude {
			err := RunCommand("git", []string{"reset", file}, workingDirectory, true)
			if err != nil {
				return
			}
		}
	}
}

func (t *ACGit) Commit(commit, workingDirectory string) {
	err := RunCommand("git", []string{"commit", "-m", commit}, workingDirectory, true)
	if err != nil {
		return
	}
}

func (t *ACGit) Push(branch, workingDirectory string, force bool, newBranch bool) error {
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

	return RunCommand("git", args, workingDirectory, true)
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

func (t *ACGit) Clone(url, workingDirectory string) error {
	err := RunCommand("git", []string{"clone", url}, workingDirectory, true)
	if err != nil {
		return err
	}

	return nil
}
