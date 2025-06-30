package dto

import (
	"backend/domain/model"
	"time"
)

type Group struct {
	ID          string        `gorm:"type:uuid;primaryKey"`
	Name        string        `gorm:"not null;size:255"`
	Description *string       `gorm:"type:text"`
	Members     []GroupMember `gorm:"foreignKey:GroupID"`
	CreatedAt   time.Time     `gorm:"autoCreateTime"`
	UpdatedAt   time.Time     `gorm:"autoUpdateTime"`
}

func (g Group) ToModel() model.Group {
	description := ""
	if g.Description != nil {
		description = *g.Description
	}

	// DTOからドメインモデルに変換
	members := make([]model.GroupMember, len(g.Members))
	for i, memberDTO := range g.Members {
		members[i] = memberDTO.ToModel()
	}

	return model.Group{
		ID:          model.GroupID(g.ID),
		Name:        g.Name,
		Description: description,
		Members:     members,
		CreatedAt:   g.CreatedAt,
	}
}

func ToGroupDto(m model.Group) Group {
	var description *string
	if m.Description != "" {
		description = &m.Description
	}

	// ドメインモデルからDTOに変換
	members := make([]GroupMember, len(m.Members))
	for i, member := range m.Members {
		members[i] = ToGroupMemberDto(member)
	}

	return Group{
		ID:          string(m.ID),
		Name:        m.Name,
		Description: description,
		Members:     members,
		CreatedAt:   m.CreatedAt,
	}
}
