package usecase

import (
	"backend/core"
	"backend/domain/model"
	"backend/domain/repository"
	"context"
	"errors"
)

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
		ID     string
		Name   string
		Role   string
		Status string
	}

	getGroupMembersInteractor struct {
		accountRepo repository.Account
		groupRepo   repository.Group
	}
)

func (g *getGroupMembersInteractor) Execute(ctx context.Context, in GetGroupMembersInput) (GetGroupMembersOutput, error) {
	account, err := g.accountRepo.FindByUserID(ctx, in.UserID)
	if err != nil {
		return GetGroupMembersOutput{}, err
	}

	groupID, err := model.ParseGroupID(in.GroupID)
	if err != nil {
		return GetGroupMembersOutput{}, err
	}

	group, err := g.groupRepo.GetByID(ctx, groupID)
	if err != nil {
		return GetGroupMembersOutput{}, err
	}

	if !group.IsMember(account.ID) {
		return GetGroupMembersOutput{}, core.NewAppError(core.ErrForbidden, errors.New("not a member of the group")).
			WithMessage("このグループのメンバーではありません")
	}

	accountIDs := g.getAccountIDs(group.Members)

	accounts, err := g.accountRepo.FindByIDs(ctx, accountIDs)
	if err != nil {
		return GetGroupMembersOutput{}, err
	}

	return g.makeOutput(group.Members, accounts), nil
}

// メンバーのアカウントIDを取得
func (g *getGroupMembersInteractor) getAccountIDs(members []model.GroupMember) []model.AccountID {
	accountIDs := make([]model.AccountID, 0, len(members))
	for _, member := range members {
		accountIDs = append(accountIDs, member.AccountID)
	}
	return accountIDs
}

// メンバーとアカウントをマッピングしてOutputを作成
func (g *getGroupMembersInteractor) makeOutput(members []model.GroupMember, accounts []model.Account) GetGroupMembersOutput {
	outputs := make([]Member, 0, len(members))

	// メンバーのIDをキーにして、アカウントをマッピング
	accountMap := make(map[model.AccountID]model.Account)
	for _, account := range accounts {
		accountMap[account.ID] = account
	}

	for _, member := range members {
		account, ok := accountMap[member.AccountID]
		if !ok {
			continue
		}

		outputs = append(outputs, Member{
			ID:     member.ID.String(),
			Name:   account.Name,
			Role:   member.Role.String(),
			Status: member.Status.String(),
		})
	}

	return GetGroupMembersOutput{
		Members: outputs,
	}
}

func NewGetGroupMembers(a repository.Account, g repository.Group) GetGroupMembers {
	return &getGroupMembersInteractor{
		accountRepo: a,
		groupRepo:   g,
	}
}
