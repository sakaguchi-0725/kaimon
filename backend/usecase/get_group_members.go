package usecase

import "context"

type (
	GetGroupMembers interface {
		Execute(ctx context.Context, in GetGroupMembersInput) (GetGroupMembersOutput, error)
	}

	GetGroupMembersInput struct {
		UserID  string
		GroupID string
	}

	GetGroupMembersOutput struct {
		Members []Member
	}

	Member struct {
		ID       string
		Name     string
		Role     string
		Status   string
		JoinedAt string
	}

	getGroupMembersInteractor struct{}
)

func (g *getGroupMembersInteractor) Execute(ctx context.Context, in GetGroupMembersInput) (GetGroupMembersOutput, error) {
	// TODO: implement get group members logic
	return GetGroupMembersOutput{
		Members: []Member{
			{
				ID:   "1",
				Name: "メンバー1",
				Role: "admin",
			},
			{
				ID:   "2",
				Name: "メンバー2",
				Role: "member",
			},
		},
	}, nil
}

func NewGetGroupMembers() GetGroupMembers {
	return &getGroupMembersInteractor{}
}
