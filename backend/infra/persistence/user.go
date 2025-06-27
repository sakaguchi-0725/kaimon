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
		return model.User{}, ErrInvalidInput
	}

	var userDTO dto.User
	err := u.conn.WithContext(ctx).Where("id = ?", id).First(&userDTO).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.User{}, ErrRecordNotFound
		}
		return model.User{}, err
	}

	return userDTO.ToModel(), nil
}

// FindByUID implements repository.User.
func (u *userPersistence) FindByUID(ctx context.Context, uid string) (*model.User, error) {
	if uid == "" {
		return nil, ErrInvalidInput
	}

	var userDTO dto.User
	err := u.conn.WithContext(ctx).Where("uid = ?", uid).First(&userDTO).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	user := userDTO.ToModel()
	return &user, nil
}

// Store implements repository.User.
func (u *userPersistence) Store(ctx context.Context, user *model.User) error {
	if user == nil {
		return ErrInvalidInput
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
