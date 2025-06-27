//go:generate mockgen -source=login.go -destination=../test/mock/usecase/login_mock.go -package=mock
package usecase

import (
	"backend/domain/repository"
	"context"
	"fmt"
)

type (
	Login interface {
		Execute(ctx context.Context, in LoginInput) error
	}

	LoginInput struct {
		Email    string
		Password string
	}

	loginInteractor struct{
		authenticator repository.Authenticator
		user          repository.User
	}
)

func (l *loginInteractor) Execute(ctx context.Context, in LoginInput) error {
	// Firebase AuthでEmail/Passwordログインを実行
	uid, err := l.authenticator.SignInWithEmailAndPassword(in.Email, in.Password)
	if err != nil {
		return fmt.Errorf("認証に失敗しました: %w", err)
	}

	// ユーザーが存在するか確認
	_, err = l.user.FindByUID(ctx, uid)
	if err != nil {
		return fmt.Errorf("ユーザーが見つかりません: %w", err)
	}

	return nil
}

func NewLogin(auth repository.Authenticator, user repository.User) Login {
	return &loginInteractor{
		authenticator: auth,
		user:          user,
	}
}
