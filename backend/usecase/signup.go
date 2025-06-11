package usecase

import "context"

type (
	SignUp interface {
		Execute(ctx context.Context, in SignUpInput) error
	}

	SignUpInput struct {
		Name     string
		Email    string
		Password string
	}

	signUpImpl struct{}
)

func (s *signUpImpl) Execute(ctx context.Context, in SignUpInput) error {
	// TODO: Implement sign up logic
	return nil
}

func NewSignUp() SignUp {
	return &signUpImpl{}
}
