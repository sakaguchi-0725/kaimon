//go:generate mockgen -source=get_joined_groups.go -destination=../test/mock/usecase/get_joined_groups_mock.go -package=mock
package usecase

import (
	"backend/domain/model"
	"backend/domain/repository"
	"context"
)

type (
	GetJoinedGroups interface {
		Execute(ctx context.Context, userID string) ([]GetJoinedGroupOutput, error)
	}

	GetJoinedGroupOutput struct {
		ID          string
		Name        string
		MemberCount int
	}

	getJoinedGroupsInteractor struct {
		account     repository.Account
		groupMember repository.GroupMember
		group       repository.Group
	}
)

func (g *getJoinedGroupsInteractor) Execute(ctx context.Context, userID string) ([]GetJoinedGroupOutput, error) {
	// ユーザーIDからアカウント情報を取得
	account, err := g.account.FindByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	// アカウントが参加しているグループメンバーシップを取得
	members, err := g.groupMember.FindByAccountID(ctx, account.ID)
	if err != nil {
		return nil, err
	}

	if len(members) == 0 {
		return []GetJoinedGroupOutput{}, nil
	}

	// グループIDを抽出
	groupIDs := make([]model.GroupID, len(members))
	for i, member := range members {
		groupIDs[i] = member.GroupID
	}

	// グループ情報を取得
	groups, err := g.group.FindByIDs(ctx, groupIDs)
	if err != nil {
		return nil, err
	}

	// レスポンス用の構造体に変換
	result := make([]GetJoinedGroupOutput, len(groups))
	for i, group := range groups {
		// 各グループのメンバー数を取得
		memberCount, err := g.groupMember.CountByGroupID(ctx, group.ID)
		if err != nil {
			return nil, err
		}

		result[i] = GetJoinedGroupOutput{
			ID:          group.ID.String(),
			Name:        group.Name,
			MemberCount: memberCount,
		}
	}

	return result, nil
}

func NewGetJoinedGroups(account repository.Account, groupMember repository.GroupMember, group repository.Group) GetJoinedGroups {
	return &getJoinedGroupsInteractor{
		account:     account,
		groupMember: groupMember,
		group:       group,
	}
}
