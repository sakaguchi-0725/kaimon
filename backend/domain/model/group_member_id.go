package model

import (
	"backend/core"
	"fmt"

	"github.com/google/uuid"
)

type GroupMemberID string

func NewGroupMemberID() GroupMemberID {
	return GroupMemberID(uuid.New().String())
}

func ParseGroupMemberID(s string) (GroupMemberID, error) {
	id, err := uuid.Parse(s)
	if err != nil {
		return "", core.NewInvalidError(fmt.Errorf("invalid group member id: %s", s))
	}

	return GroupMemberID(id.String()), nil
}

func (id GroupMemberID) String() string {
	return string(id)
}
