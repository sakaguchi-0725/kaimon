package dto

import (
	"backend/domain/model"
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID              uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserUID         string    `gorm:"not null;unique;constraint:fk_accounts_user_uid,OnDelete:CASCADE"`
	Name            string    `gorm:"not null"`
	ProfileImageURL *string   `gorm:"size:500"`
	IsPremium       bool      `gorm:"default:false"`
	CreatedAt       time.Time `gorm:"autoCreateTime"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime"`

	User User `gorm:"foreignKey:UserUID;references:ID"`
}

// ToModel はDTOからドメインモデルに変換する
func (a Account) ToModel() model.Account {
	return model.Account{
		ID:     model.AccountID(a.ID.String()),
		UserID: a.UserUID,
		Name:   a.Name,
	}
}

// ToAccountDto はドメインモデルからDTOに変換する
func ToAccountDto(m model.Account) (Account, error) {
	id, err := uuid.Parse(string(m.ID))
	if err != nil {
		return Account{}, err
	}

	return Account{
		ID:      id,
		UserUID: m.UserID,
		Name:    m.Name,
	}, nil
}
