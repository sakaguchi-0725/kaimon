//go:generate mockgen -source=get_group.go -destination=../test/mock/usecase/get_group_mock.go -package=mock
package usecase

import (
	"backend/core"
	"backend/domain/model"
	"backend/domain/repository"
	"context"
	"errors"
)

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
		Members     []Member
	}

	Member struct {
		ID     string
		Name   string
		Role   string
		Status string
	}

	getGroupInteractor struct {
		accountRepo repository.Account
		groupRepo   repository.Group
	}
)

func (g *getGroupInteractor) Execute(ctx context.Context, in GetGroupInput) (GetGroupOutput, error) {
	account, err := g.accountRepo.FindByUserID(ctx, in.UserID)
	if err != nil {
		return GetGroupOutput{}, err
	}

	groupID, err := model.ParseGroupID(in.GroupID)
	if err != nil {
		return GetGroupOutput{}, err
	}

	group, err := g.groupRepo.GetByID(ctx, groupID)
	if err != nil {
		return GetGroupOutput{}, err
	}

	if !group.IsMember(account.ID) {
		return GetGroupOutput{}, core.NewAppError(core.ErrForbidden, errors.New("not a member of the group")).
			WithMessage("このグループのメンバーではありません")
	}

	accountIDs := g.getAccountIDs(group.Members)

	accounts, err := g.accountRepo.FindByIDs(ctx, accountIDs)
	if err != nil {
		return GetGroupOutput{}, err
	}

	return g.makeOutput(group, accounts), nil
}

// メンバーのアカウントIDを取得
func (g *getGroupInteractor) getAccountIDs(members []model.GroupMember) []model.AccountID {
	accountIDs := make([]model.AccountID, 0, len(members))
	for _, member := range members {
		accountIDs = append(accountIDs, member.AccountID)
	}
	return accountIDs
}

// グループとメンバー情報をマッピングしてOutputを作成
func (g *getGroupInteractor) makeOutput(group model.Group, accounts []model.Account) GetGroupOutput {
	members := make([]Member, 0, len(group.Members))

	// メンバーのIDをキーにして、アカウントをマッピング
	accountMap := make(map[model.AccountID]model.Account)
	for _, account := range accounts {
		accountMap[account.ID] = account
	}

	for _, member := range group.Members {
		account, ok := accountMap[member.AccountID]
		if !ok {
			continue
		}

		members = append(members, Member{
			ID:     member.ID.String(),
			Name:   account.Name,
			Role:   member.Role.String(),
			Status: member.Status.String(),
		})
	}

	return GetGroupOutput{
		ID:          group.ID.String(),
		Name:        group.Name,
		Description: group.Description,
		CreatedAt:   group.CreatedAt.Format("2006-01-02T15:04:05Z"),
		Members:     members,
	}
}

func NewGetGroup(a repository.Account, g repository.Group) GetGroup {
	return &getGroupInteractor{
		accountRepo: a,
		groupRepo:   g,
	}
}
