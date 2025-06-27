package persistence

import (
	"backend/domain/model"
	"backend/domain/repository"
	"backend/infra/db"
	"backend/infra/dto"
	"context"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type accountPersistence struct {
	conn *db.Conn
}

// FindByID implements repository.Account.
func (a *accountPersistence) FindByID(ctx context.Context, id model.AccountID) (model.Account, error) {
	if id == "" {
		return model.Account{}, errors.New("id is required")
	}

	parsedID, err := uuid.Parse(string(id))
	if err != nil {
		return model.Account{}, errors.New("invalid account id format")
	}

	var accountDTO dto.Account
	err = a.conn.WithContext(ctx).Where("id = ?", parsedID).First(&accountDTO).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Account{}, errors.New("account not found")
		}
		return model.Account{}, err
	}

	return accountDTO.ToModel(), nil
}

// Store implements repository.Account.
func (a *accountPersistence) Store(ctx context.Context, acc *model.Account) error {
	if acc == nil {
		return errors.New("account is nil")
	}

	accountDTO, err := dto.ToAccountDto(*acc)
	if err != nil {
		return err
	}

	// UPSERTの実装（存在すればUpdate、なければInsert）
	err = a.conn.WithContext(ctx).Save(&accountDTO).Error
	if err != nil {
		return err
	}

	return nil
}

func NewAccount(c *db.Conn) repository.Account {
	return &accountPersistence{conn: c}
}
