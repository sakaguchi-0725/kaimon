package dto

import (
	"backend/domain/model"
	"time"
)

type Account struct {
	ID              string    `gorm:"type:uuid;primaryKey"`
	UserUID         string    `gorm:"not null;unique;constraint:fk_accounts_user_uid,OnDelete:CASCADE"`
	Name            string    `gorm:"not null"`
	ProfileImageURL *string   `gorm:"size:500"`
	IsPremium       bool      `gorm:"default:false"`
	CreatedAt       time.Time `gorm:"autoCreateTime"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime"`

	User User `gorm:"foreignKey:UserUID;references:ID"`
}

func (a Account) ToModel() model.Account {
	return model.Account{
		ID:     model.AccountID(a.ID),
		UserID: a.UserUID,
		Name:   a.Name,
	}
}

func ToAccountDto(m model.Account) Account {
	return Account{
		ID:      m.ID.String(),
		UserUID: m.UserID,
		Name:    m.Name,
	}
}
