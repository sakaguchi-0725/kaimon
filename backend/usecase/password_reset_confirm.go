package usecase

import "context"

type (
	ResetPasswordConfirm interface {
		Execute(ctx context.Context, in ResetPasswordConfirmInput) error
	}

	ResetPasswordConfirmInput struct {
		Email            string
		ConfirmationCode string
	}

	resetPasswordConfirmInteractor struct{}
)

func (r *resetPasswordConfirmInteractor) Execute(ctx context.Context, in ResetPasswordConfirmInput) error {
	// TODO: Implement password reset confirm logic
	return nil
}

func NewResetPasswordConfirm() ResetPasswordConfirm {
	return &resetPasswordConfirmInteractor{}
}
