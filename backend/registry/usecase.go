package registry

import "backend/usecase"

type UseCase struct {
	SignUp        usecase.SignUp
	SignUpConfirm usecase.SignUpConfirm
}

func NewUseCase() UseCase {
	return UseCase{
		SignUp:        usecase.NewSignUp(),
		SignUpConfirm: usecase.NewSignUpConfirm(),
	}
}
