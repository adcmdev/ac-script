package in

type IGit interface {
	Add(files string, workingDirectory string, filesToExclude []string)
	Commit(commit string, workingDirectory string)
	Push(branch string, workingDirectory string, force bool, newBranch bool) bool
}
