package dto

import (
	"backend/domain/model"
	"time"
)

type User struct {
	ID        string    `gorm:"primaryKey;size:255"`
	Email     string    `gorm:"not null;unique;size:255"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	Account *Account `gorm:"foreignKey:UserUID;references:ID"`
}

// ToModel はDTOからドメインモデルに変換する
func (u User) ToModel() model.User {
	return model.User{
		ID:    u.ID,
		Email: u.Email,
	}
}

// ToUserDto はドメインモデルからDTOに変換する
func ToUserDto(m model.User) User {
	return User{
		ID:    m.ID,
		Email: m.Email,
	}
}
