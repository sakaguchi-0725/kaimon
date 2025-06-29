package model

type MemberStatus string

const (
	MemberStatusActive  MemberStatus = "active"
	MemberStatusPending MemberStatus = "pending"
)

func (s MemberStatus) String() string {
	return string(s)
}
