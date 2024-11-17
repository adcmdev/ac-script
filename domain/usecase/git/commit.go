package git

import "ac/domain/port/in"

type UsecaseService struct {
	git in.IGit
}

func NewUsecaseService(git in.IGit) in.ICommitUsecase {
	return &UsecaseService{
		git: git,
	}
}

func (u *UsecaseService) Commit(message string) error {
	u.git.Add(".", "", []string{})
	u.git.Commit(message, "")
	u.git.Push("", "", false, false)

	return nil
}
