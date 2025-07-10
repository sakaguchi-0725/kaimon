//go:generate mockgen -source=get_account.go -destination=../test/mock/usecase/get_account_mock.go -package=mock
package usecase

import (
	"backend/domain/repository"
	"context"
)

type (
	GetAccount interface {
		Execute(ctx context.Context, input GetAccountInput) (GetAccountOutput, error)
	}

	GetAccountInput struct {
		UserID string
	}

	GetAccountOutput struct {
		ID   string
		Name string
	}

	getAccountInteractor struct {
		accountRepo repository.Account
	}
)

func (g *getAccountInteractor) Execute(ctx context.Context, input GetAccountInput) (GetAccountOutput, error) {
	account, err := g.accountRepo.FindByUserID(ctx, input.UserID)
	if err != nil {
		return GetAccountOutput{}, err
	}

	return GetAccountOutput{
		ID:   account.ID.String(),
		Name: account.Name,
	}, nil
}

func NewGetAccount(a repository.Account) GetAccount {
	return &getAccountInteractor{
		accountRepo: a,
	}
}
