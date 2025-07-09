package registry

import "backend/usecase"

type UseCase struct {
	SignUp                 usecase.SignUp
	ResendConfirmationCode usecase.ResendConfirmationCode
	VerifyToken            usecase.VerifyToken
	GetJoinedGroups        usecase.GetJoinedGroups
	CreateGroup            usecase.CreateGroup
	GetGroup               usecase.GetGroup
	GetShoppingItems       usecase.GetShoppingItems
	GroupInvitation        usecase.GroupInvitation
}

func NewUseCase(repo *Repository) UseCase {
	return UseCase{
		SignUp:                 usecase.NewSignUp(repo.Account, repo.User, repo.Transaction),
		ResendConfirmationCode: usecase.NewResendConfirmationCode(),
		VerifyToken:            usecase.NewVerifyToken(repo.Authenticator),
		GetJoinedGroups:        usecase.NewGetJoinedGroups(repo.Account, repo.GroupMember, repo.Group),
		CreateGroup:            usecase.NewCreateGroup(repo.Account, repo.Group),
		GetGroup:               usecase.NewGetGroup(repo.Account, repo.Group),
		GetShoppingItems:       usecase.NewGetShoppingItems(),
		GroupInvitation:        usecase.NewGroupInvitation(repo.Account, repo.Group),
	}
}
