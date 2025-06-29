package persistence

import (
	"backend/domain/model"
	"backend/domain/repository"
	"backend/infra/db"
	"backend/infra/dto"
	"context"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type groupMemberPersistence struct {
	conn *db.Conn
}

func (g *groupMemberPersistence) FindByAccountID(ctx context.Context, accountID model.AccountID) ([]model.GroupMember, error) {
	var memberDTOs []dto.GroupMember
	err := g.conn.WithContext(ctx).Where("account_id = ?", accountID.String()).Find(&memberDTOs).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []model.GroupMember{}, nil
		}
		return []model.GroupMember{}, err
	}

	// DTOからドメインモデルに変換
	members := make([]model.GroupMember, len(memberDTOs))
	for i, memberDTO := range memberDTOs {
		members[i] = memberDTO.ToModel()
	}

	return members, nil
}

func (g *groupMemberPersistence) CountByGroupID(ctx context.Context, groupID model.GroupID) (int, error) {
	// UUIDの妥当性チェック
	_, err := uuid.Parse(string(groupID))
	if err != nil {
		return 0, ErrInvalidInput
	}

	var count int64
	err = g.conn.WithContext(ctx).Model(&dto.GroupMember{}).Where("group_id = ?", string(groupID)).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return int(count), nil
}

func NewGroupMember(c *db.Conn) repository.GroupMember {
	return &groupMemberPersistence{conn: c}
}
