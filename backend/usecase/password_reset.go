package usecase

import "context"

type (
	ResetPassword interface {
		Execute(ctx context.Context, in ResetPasswordInput) error
	}

	ResetPasswordInput struct {
		Email string
	}

	resetPasswordInteractor struct{}
)

func (r *resetPasswordInteractor) Execute(ctx context.Context, in ResetPasswordInput) error {
	// TODO: Implement reset password logic
	return nil
}

func NewResetPassword() ResetPassword {
	return &resetPasswordInteractor{}
}
