package persistence

import (
	"backend/domain/model"
	"backend/domain/repository"
	"backend/infra/db"
	"backend/infra/dto"
	"context"
	"errors"

	"gorm.io/gorm"
)

type userPersistence struct {
	conn *db.Conn
}

// FindByID implements repository.User.
func (u *userPersistence) FindByID(ctx context.Context, id string) (model.User, error) {
	if id == "" {
		return model.User{}, errors.New("id is required")
	}

	var userDTO dto.User
	err := u.conn.WithContext(ctx).Where("id = ?", id).First(&userDTO).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.User{}, errors.New("user not found")
		}
		return model.User{}, err
	}

	return userDTO.ToModel(), nil
}

// Store implements repository.User.
func (u *userPersistence) Store(ctx context.Context, user *model.User) error {
	if user == nil {
		return errors.New("user is nil")
	}

	userDTO := dto.ToUserDto(*user)

	// UPSERTの実装（存在すればUpdate、なければInsert）
	err := u.conn.WithContext(ctx).Save(&userDTO).Error
	if err != nil {
		return err
	}

	return nil
}

func NewUser(c *db.Conn) repository.User {
	return &userPersistence{conn: c}
}
