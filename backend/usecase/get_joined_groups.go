package usecase

import "context"

type (
	GetJoinedGroups interface {
		Execute(ctx context.Context, userID string) ([]GetJoinedGroupOutput, error)
	}

	GetJoinedGroupOutput struct {
		ID          string
		Name        string
		Description string
	}

	getJoinedGroupsInteractor struct{}
)

func (g *getJoinedGroupsInteractor) Execute(ctx context.Context, userID string) ([]GetJoinedGroupOutput, error) {
	return []GetJoinedGroupOutput{
		{
			ID:          "1",
			Name:        "テストグループ",
			Description: "モックのグループです",
		},
	}, nil
}

func NewGetJoinedGroups() GetJoinedGroups {
	return &getJoinedGroupsInteractor{}
}
