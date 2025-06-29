package dto

import (
	"backend/domain/model"
	"time"
)

type GroupMember struct {
	ID        string    `gorm:"type:uuid;primaryKey"`
	GroupID   string    `gorm:"type:uuid;not null;constraint:fk_members_group_id,OnDelete:CASCADE"`
	AccountID string    `gorm:"type:uuid;not null;constraint:fk_members_account_id,OnDelete:CASCADE"`
	Role      string    `gorm:"not null;size:255"`
	Status    string    `gorm:"not null;size:255"`
	JoinedAt  time.Time `gorm:"autoCreateTime"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	Group   Group   `gorm:"foreignKey:GroupID;references:ID"`
	Account Account `gorm:"foreignKey:AccountID;references:ID"`
}

func (gm GroupMember) ToModel() model.GroupMember {
	return model.GroupMember{
		ID:        model.GroupMemberID(gm.ID),
		GroupID:   model.GroupID(gm.GroupID),
		AccountID: model.AccountID(gm.AccountID),
		Role:      model.MemberRole(gm.Role),
		Status:    model.MemberStatus(gm.Status),
		JoinedAt:  gm.JoinedAt,
	}
}

func ToGroupMemberDto(m model.GroupMember) GroupMember {
	return GroupMember{
		ID:        m.ID.String(),
		GroupID:   m.GroupID.String(),
		AccountID: m.AccountID.String(),
		Role:      string(m.Role),
		Status:    string(m.Status),
		JoinedAt:  m.JoinedAt,
	}
}
