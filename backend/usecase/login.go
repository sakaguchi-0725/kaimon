package usecase

import (
	"context"
)

type (
	Login interface {
		Execute(ctx context.Context, in LoginInput) error
	}

	LoginInput struct {
		Email    string
		Password string
	}

	loginInteractor struct{}
)

func (l *loginInteractor) Execute(ctx context.Context, in LoginInput) error {
	// TODO: Implement login logic
	return nil
}

func NewLogin() Login {
	return &loginInteractor{}
}
