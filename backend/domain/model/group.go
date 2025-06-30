package model

import (
	"backend/core"
	"fmt"
	"time"
)

type Group struct {
	ID          GroupID
	Name        string
	Description string
	Members     []GroupMember
	CreatedAt   time.Time
}

func NewGroup(id GroupID, name string, description string) (Group, error) {
	if name == "" {
		return Group{}, core.NewInvalidError(fmt.Errorf("name is required"))
	}

	return Group{
		ID:          id,
		Name:        name,
		Description: description,
		CreatedAt:   core.NowJST(),
	}, nil
}

// アクティブなメンバーかどうかを判定する
func (g Group) IsMember(accountID AccountID) bool {
	for _, member := range g.Members {
		if member.AccountID == accountID && member.Status == MemberStatusActive {
			return true
		}
	}
	return false
}
