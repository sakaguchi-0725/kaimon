package registry

import "backend/usecase"

type UseCase struct {
	SignUp        usecase.SignUp
	SignUpConfirm usecase.SignUpConfirm
	Login         usecase.Login
}

func NewUseCase() UseCase {
	return UseCase{
		SignUp:        usecase.NewSignUp(),
		SignUpConfirm: usecase.NewSignUpConfirm(),
		Login:         usecase.NewLogin(),
	}
}
