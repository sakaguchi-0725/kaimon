package dto

import (
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
