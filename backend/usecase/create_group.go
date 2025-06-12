package usecase

import "context"

type (
	CreateGroup interface {
		Execute(ctx context.Context, in CreateGroupInput) error
	}

	CreateGroupInput struct {
		UserID      string
		Name        string
		Description string
	}

	createGroupInteractor struct{}
)

func (r *createGroupInteractor) Execute(ctx context.Context, in CreateGroupInput) error {
	// TODO: Implement create group logic
	return nil
}

func NewCreateGroup() CreateGroup {
	return &createGroupInteractor{}
}
