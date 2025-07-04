package usecase

import (
	"backend/domain/model"
	"backend/domain/repository"
	"context"
)

type (
	CreateGroup interface {
		Execute(ctx context.Context, in CreateGroupInput) error
	}

	CreateGroupInput struct {
		UserID      string
		Name        string
		Description string
	}

	createGroupInteractor struct {
		accountRepo repository.Account
		groupRepo   repository.Group
	}
)

func (r *createGroupInteractor) Execute(ctx context.Context, in CreateGroupInput) error {
	// アカウント取得
	account, err := r.accountRepo.FindByUserID(ctx, in.UserID)
	if err != nil {
		return err
	}

	// グループメンバー作成（作成者をAdmin）
	groupID := model.NewGroupID()
	member, err := model.NewGroupMember(
		model.NewGroupMemberID(),
		groupID,
		account.ID,
	)
	if err != nil {
		return err
	}

	// グループ作成
	group, err := model.NewGroup(groupID, in.Name, in.Description, []model.GroupMember{member})
	if err != nil {
		return err
	}

	// データ永続化
	return r.groupRepo.Store(ctx, &group)
}

func NewCreateGroup(accountRepo repository.Account, groupRepo repository.Group) CreateGroup {
	return &createGroupInteractor{
		accountRepo: accountRepo,
		groupRepo:   groupRepo,
	}
}
