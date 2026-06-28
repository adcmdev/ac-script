package in

type ICheckoutUsecase interface {
	Checkout(branch string) error
}
