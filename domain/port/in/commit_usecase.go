package in

type ICommitUsecase interface {
	Commit(message string) error
}
