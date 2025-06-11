package registry

import "backend/usecase"

type UseCase struct {
	SignUp usecase.SignUp
}

func NewUseCase() UseCase {
	return UseCase{
		SignUp: usecase.NewSignUp(),
	}
}
