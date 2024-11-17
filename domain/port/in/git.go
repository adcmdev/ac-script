package in

type IGit interface {
	Init(repoName, workingDirectory string) error
	Add(files string, workingDirectory string, filesToExclude []string)
	Commit(commit string, workingDirectory string)
	Push(branch string, workingDirectory string, force bool, newBranch bool) error
	Clone(url string, workingDirectory string) error
}
