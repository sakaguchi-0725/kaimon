package usecase

import "context"

type (
	GetGroup interface {
		Execute(ctx context.Context, in GetGroupInput) (GetGroupOutput, error)
	}

	GetGroupInput struct {
		UserID  string
		GroupID string
	}

	GetGroupOutput struct {
		ID          string
		Name        string
		Description string
		CreatedAt   string
	}

	getGroupInteractor struct{}
)

func (r *getGroupInteractor) Execute(ctx context.Context, in GetGroupInput) (GetGroupOutput, error) {
	// TODO: Implement get group logic
	return GetGroupOutput{
		ID:          "group-id",
		Name:        "テストグループ",
		Description: "テストグループの説明",
		CreatedAt:   "2021-01-01T00:00:00Z",
	}, nil
}

func NewGetGroup() GetGroup {
	return &getGroupInteractor{}
}
