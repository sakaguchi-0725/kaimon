package usecase

import "context"

type (
	ResendConfirmationCode interface {
		Execute(ctx context.Context, in ResendConfirmationCodeInput) error
	}

	ResendConfirmationCodeInput struct {
		Email string
	}

	resendConfirmationCodeInteractor struct{}
)

func (r *resendConfirmationCodeInteractor) Execute(ctx context.Context, in ResendConfirmationCodeInput) error {
	// TODO: Implement resend confirmation code logic
	return nil
}

func NewResendConfirmationCode() ResendConfirmationCode {
	return &resendConfirmationCodeInteractor{}
}
