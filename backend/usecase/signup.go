//go:generate mockgen -source=signup.go -destination=../test/mock/usecase/signup_mock.go -package=mock
package usecase

import (
	"backend/domain/model"
	"backend/domain/repository"
	"context"
)

type (
	SignUp interface {
		Execute(ctx context.Context, in SignUpInput) error
	}

	SignUpInput struct {
		UID             string
		Email           string
		Name            string
		ProfileImageURL string
	}

	signUpInteractor struct {
		account repository.Account
		user    repository.User
		tx      repository.Transaction
	}
)

func (s *signUpInteractor) Execute(ctx context.Context, in SignUpInput) error {
	// トランザクション内でユーザーとアカウントを作成
	err := s.tx.WithTx(ctx, func(txCtx context.Context) error {
		user, err := model.NewUser(in.UID, in.Email)
		if err != nil {
			return err
		}

		if err := s.user.Store(txCtx, &user); err != nil {
			return err
		}

		accountID := model.NewAccountID()
		account, err := model.NewAccount(accountID, user.ID, in.Name, in.ProfileImageURL)
		if err != nil {
			return err
		}

		if err := s.account.Store(txCtx, &account); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func NewSignUp(acc repository.Account, user repository.User, tx repository.Transaction) SignUp {
	return &signUpInteractor{
		account: acc,
		user:    user,
		tx:      tx,
	}
}
