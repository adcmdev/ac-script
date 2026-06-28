package git

import (
	"ac/domain/port/in"
)

type CheckoutUsecaseService struct {
	git in.IGit
}

func NewCheckoutUsecaseService(git in.IGit) in.ICheckoutUsecase {
	return &CheckoutUsecaseService{
		git: git,
	}
}

func (u *CheckoutUsecaseService) Checkout(branch string) error {
	return u.git.Checkout(branch, "")
}
