//go:generate mockgen -source=signup.go -destination=../test/mock/usecase/signup_mock.go -package=mock
package usecase

import (
	"backend/domain/model"
	"backend/domain/repository"
	"context"
	"fmt"
)

type (
	SignUp interface {
		Execute(ctx context.Context, in SignUpInput) error
	}

	SignUpInput struct {
		IDToken string
		Name    string
	}

	signUpInteractor struct {
		authenticator repository.Authenticator
		account       repository.Account
		user          repository.User
		tx            repository.Transaction
	}
)

func (s *signUpInteractor) Execute(ctx context.Context, in SignUpInput) error {
	// Firebase IDトークンを検証してユーザー情報を取得
	uid, email, err := s.authenticator.VerifyToken(in.IDToken)
	if err != nil {
		return fmt.Errorf("認証に失敗しました: %w", err)
	}

	// トランザクション内でユーザーとアカウントを作成
	err = s.tx.WithTx(ctx, func(txCtx context.Context) error {
		user, err := model.NewUser(uid, email)
		if err != nil {
			return err
		}

		if err := s.user.Store(txCtx, &user); err != nil {
			return err
		}

		accountID := model.NewAccountID()
		account, err := model.NewAccount(accountID, uid, in.Name)
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

func NewSignUp(auth repository.Authenticator, acc repository.Account, user repository.User, tx repository.Transaction) SignUp {
	return &signUpInteractor{
		authenticator: auth,
		account:       acc,
		user:          user,
		tx:            tx,
	}
}
