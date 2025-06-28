package registry

import "backend/usecase"

type UseCase struct {
	SignUp                 usecase.SignUp
	ResendConfirmationCode usecase.ResendConfirmationCode
	ResetPassword          usecase.ResetPassword
	ResetPasswordConfirm   usecase.ResetPasswordConfirm
	VerifyToken            usecase.VerifyToken
	GetJoinedGroups        usecase.GetJoinedGroups
	CreateGroup            usecase.CreateGroup
	GetGroup               usecase.GetGroup
	GetGroupMembers        usecase.GetGroupMembers
	GetShoppingItems       usecase.GetShoppingItems
}

func NewUseCase(repo *Repository) UseCase {
	return UseCase{
		SignUp:                 usecase.NewSignUp(repo.Authenticator, repo.Account, repo.User, repo.Transaction),
		ResendConfirmationCode: usecase.NewResendConfirmationCode(),
		ResetPassword:          usecase.NewResetPassword(),
		ResetPasswordConfirm:   usecase.NewResetPasswordConfirm(),
		VerifyToken:            usecase.NewVerifyToken(repo.Authenticator),
		GetJoinedGroups:        usecase.NewGetJoinedGroups(),
		CreateGroup:            usecase.NewCreateGroup(),
		GetGroup:               usecase.NewGetGroup(),
		GetGroupMembers:        usecase.NewGetGroupMembers(),
		GetShoppingItems:       usecase.NewGetShoppingItems(),
	}
}
