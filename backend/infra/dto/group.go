package dto

import (
	"backend/domain/model"
	"time"
)

type Group struct {
	ID          string    `gorm:"type:uuid;primaryKey"`
	Name        string    `gorm:"not null;size:255"`
	Description *string   `gorm:"type:text"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}

func (g Group) ToModel() model.Group {
	description := ""
	if g.Description != nil {
		description = *g.Description
	}

	return model.Group{
		ID:          model.GroupID(g.ID),
		Name:        g.Name,
		Description: description,
		CreatedAt:   g.CreatedAt,
	}
}

func ToGroupDto(m model.Group) Group {
	var description *string
	if m.Description != "" {
		description = &m.Description
	}

	return Group{
		ID:          string(m.ID),
		Name:        m.Name,
		Description: description,
		CreatedAt:   m.CreatedAt,
	}
}
