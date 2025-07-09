package model

import (
	"backend/core"
	"math/rand"
	"time"
)

type Invitation struct {
	Code      string
	GroupID   GroupID
	ExpiresAt time.Time
}

func NewInvitation(groupID GroupID) Invitation {
	return Invitation{
		Code:      generateInvitationCode(),
		GroupID:   groupID,
		ExpiresAt: core.NowJST().Add(7 * 24 * time.Hour), // 7日後に期限切れ
	}
}

func generateInvitationCode() string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const length = 8

	code := make([]byte, length)
	for i := range code {
		code[i] = charset[rand.Intn(len(charset))]
	}
	return string(code)
}
