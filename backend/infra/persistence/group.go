package persistence

import (
	"backend/domain/model"
	"backend/domain/repository"
	"backend/infra/db"
	"backend/infra/dto"
	"context"
	"errors"

	"gorm.io/gorm"
)

type groupPersistence struct {
	conn *db.Conn
}

func (g *groupPersistence) GetByID(ctx context.Context, id model.GroupID) (model.Group, error) {
	var groupDTO dto.Group
	err := g.conn.WithContext(ctx).Where("id = ?", id.String()).First(&groupDTO).Error
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
	err := g.conn.WithContext(ctx).Where("id IN ?", stringIDs).Find(&groupDTOs).Error
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

func NewGroup(c *db.Conn) repository.Group {
	return &groupPersistence{conn: c}
}
