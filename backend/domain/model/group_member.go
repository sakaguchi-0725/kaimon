package model

import (
	"backend/core"
	"time"
)

type GroupMember struct {
	ID        GroupMemberID
	GroupID   GroupID
	AccountID AccountID
	Role      MemberRole
	Status    MemberStatus
	JoinedAt  time.Time
}

// グループ作成者を管理者として追加する
func NewGroupMember(id GroupMemberID, groupID GroupID, accountID AccountID) (GroupMember, error) {
	return GroupMember{
		ID:        id,
		GroupID:   groupID,
		AccountID: accountID,
		Role:      MemberRoleAdmin,
		Status:    MemberStatusActive,
		JoinedAt:  core.NowJST(),
	}, nil
}
