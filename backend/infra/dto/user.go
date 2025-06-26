package dto

import "time"

type User struct {
	ID        string    `gorm:"primaryKey;size:255"`
	Email     string    `gorm:"not null;unique;size:255"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	Account *Account `gorm:"foreignKey:UserUID;references:ID"`
}
