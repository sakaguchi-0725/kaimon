package registry

import "backend/usecase"

type UseCase struct {
	SignUp                 usecase.SignUp
	SignUpConfirm          usecase.SignUpConfirm
	ResendConfirmationCode usecase.ResendConfirmationCode
	Login                  usecase.Login
	ResetPassword          usecase.ResetPassword
	ResetPasswordConfirm   usecase.ResetPasswordConfirm
	VerifyToken            usecase.VerifyToken
	GetJoinedGroups        usecase.GetJoinedGroups
	CreateGroup            usecase.CreateGroup
	GetGroup               usecase.GetGroup
}

func NewUseCase() UseCase {
	return UseCase{
		SignUp:                 usecase.NewSignUp(),
		SignUpConfirm:          usecase.NewSignUpConfirm(),
		ResendConfirmationCode: usecase.NewResendConfirmationCode(),
		Login:                  usecase.NewLogin(),
		ResetPassword:          usecase.NewResetPassword(),
		ResetPasswordConfirm:   usecase.NewResetPasswordConfirm(),
		VerifyToken:            usecase.NewVerifyToken(),
		GetJoinedGroups:        usecase.NewGetJoinedGroups(),
		CreateGroup:            usecase.NewCreateGroup(),
		GetGroup:               usecase.NewGetGroup(),
	}
}
