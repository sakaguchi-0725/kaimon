package model

import (
	"backend/core"
	"fmt"

	"github.com/google/uuid"
)

type GroupID string

func NewGroupID() GroupID {
	return GroupID(uuid.New().String())
}

func ParseGroupID(s string) (GroupID, error) {
	id, err := uuid.Parse(s)
	if err != nil {
		return "", core.NewInvalidError(fmt.Errorf("invalid group id: %s", s))
	}

	return GroupID(id.String()), nil
}

func (id GroupID) String() string {
	return string(id)
}
