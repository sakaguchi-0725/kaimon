package usecase

import (
	"backend/domain/repository"
	"context"
)

type (
	VerifyToken interface {
		Execute(ctx context.Context, token string) (uid, email string, err error)
	}

	verifyTokenInteractor struct {
		authenticator repository.Authenticator
	}
)

func (r *verifyTokenInteractor) Execute(ctx context.Context, token string) (uid, email string, err error) {
	uid, email, _, err = r.authenticator.VerifyToken(ctx, token)
	if err != nil {
		return "", "", err
	}

	return uid, email, nil
}

func NewVerifyToken(a repository.Authenticator) VerifyToken {
	return &verifyTokenInteractor{
		authenticator: a,
	}
}
