package usecase

import "context"

type (
	SignUpConfirm interface {
		Execute(ctx context.Context, in SignUpConfirmInput) error
	}

	SignUpConfirmInput struct {
		Email            string
		ConfirmationCode string
	}

	signUpConfirmInteractor struct{}
)

func (s *signUpConfirmInteractor) Execute(ctx context.Context, in SignUpConfirmInput) error {
	// TODO: Implement sign up confirm logic
	return nil
}

func NewSignUpConfirm() SignUpConfirm {
	return &signUpConfirmInteractor{}
}
