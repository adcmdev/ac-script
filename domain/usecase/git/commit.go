package git

import (
	"ac/domain/port/in"
	"ac/shared"
)

type UsecaseService struct {
	git in.IGit
}

func NewUsecaseService(git in.IGit) in.ICommitUsecase {
	return &UsecaseService{
		git: git,
	}
}

func (u *UsecaseService) Commit(message string) (err error) {
	err = shared.RunCommand("dart", []string{"format", "."}, "", true)
	if err != nil {
		return
	}

	u.git.Add(".", "", []string{})
	u.git.Commit(message, "")
	err = u.git.Push("", "", false, false)

	return
}
