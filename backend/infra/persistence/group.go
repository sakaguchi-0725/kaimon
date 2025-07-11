package persistence

import (
	"backend/core"
	"backend/domain/model"
	"backend/domain/repository"
	"backend/infra/db"
	"backend/infra/dto"
	"backend/infra/external"
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type groupPersistence struct {
	conn        *db.Conn
	redisClient external.RedisClient
}

func (g *groupPersistence) Invitation(ctx context.Context, invitation model.Invitation) error {
	invitationJSON, err := json.Marshal(invitation)
	if err != nil {
		return err
	}

	key := g.generateGroupInvitationKey(invitation.GroupID)
	err = g.redisClient.Set(ctx, key, invitationJSON, invitation.ExpiresAt.Sub(core.NowJST()))
	if err != nil {
		return err
	}

	return nil
}

func (g *groupPersistence) GetInvitation(ctx context.Context, groupID model.GroupID) (*model.Invitation, error) {
	key := g.generateGroupInvitationKey(groupID)
	result, err := g.redisClient.Get(ctx, key)
	if err != nil {
		return nil, err
	}

	if result == "" {
		return nil, nil // 招待コードが存在しない
	}

	var invitation model.Invitation
	err = json.Unmarshal([]byte(result), &invitation)
	if err != nil {
		return nil, err
	}

	return &invitation, nil
}

func (g *groupPersistence) Store(ctx context.Context, group *model.Group) error {
	groupDTO := dto.ToGroupDto(*group)
	return g.conn.WithContext(ctx).Create(&groupDTO).Error
}

func (g *groupPersistence) GetByID(ctx context.Context, id model.GroupID) (model.Group, error) {
	var groupDTO dto.Group
	err := g.conn.WithContext(ctx).Preload("Members").Where("id = ?", id.String()).First(&groupDTO).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Group{}, ErrRecordNotFound
		}
		return model.Group{}, err
	}

	return groupDTO.ToModel(), nil
}

func (g *groupPersistence) FindByIDs(ctx context.Context, ids []model.GroupID) ([]model.Group, error) {
	if len(ids) == 0 {
		return []model.Group{}, nil
	}

	// UUIDの文字列変換
	stringIDs := make([]string, len(ids))
	for i, id := range ids {
		stringIDs[i] = id.String()
	}

	var groupDTOs []dto.Group
	err := g.conn.WithContext(ctx).Preload("Members").Where("id IN ?", stringIDs).Find(&groupDTOs).Error
	if err != nil {
		return nil, err
	}

	// DTOからドメインモデルに変換
	groups := make([]model.Group, len(groupDTOs))
	for i, groupDTO := range groupDTOs {
		groups[i] = groupDTO.ToModel()
	}

	return groups, nil
}

func (g *groupPersistence) generateGroupInvitationKey(groupID model.GroupID) string {
	return fmt.Sprintf("group_invitation:%s", groupID.String())
}

func NewGroup(c *db.Conn, redisClient external.RedisClient) repository.Group {
	return &groupPersistence{
		conn:        c,
		redisClient: redisClient,
	}
}
