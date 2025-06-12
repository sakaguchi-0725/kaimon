package usecase

import "context"

type (
	VerifyToken interface {
		Execute(ctx context.Context, token string) (string, error)
	}

	verifyTokenInteractor struct{}
)

// Tokenを検証し、ユーザーIDを返す
func (r *verifyTokenInteractor) Execute(ctx context.Context, token string) (string, error) {
	// TODO: Implement verify token logic
	return "dummy-user-id", nil
}

func NewVerifyToken() VerifyToken {
	return &verifyTokenInteractor{}
}
