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

func (a *accountPersistence) FindByIDs(ctx context.Context, ids []model.AccountID) ([]model.Account, error) {
	if len(ids) == 0 {
		return []model.Account{}, nil
	}

	// AccountIDをstring型に変換
	stringIDs := make([]string, len(ids))
	for i, id := range ids {
		stringIDs[i] = id.String()
	}

	var accountDTOs []dto.Account
	err := a.conn.WithContext(ctx).Where("id IN ?", stringIDs).Find(&accountDTOs).Error
	if err != nil {
		return nil, err
	}

	// DTOからドメインモデルに変換
	accounts := make([]model.Account, len(accountDTOs))
	for i, accountDTO := range accountDTOs {
		accounts[i] = accountDTO.ToModel()
	}

	return accounts, nil
}

func (a *accountPersistence) FindByID(ctx context.Context, id model.AccountID) (model.Account, error) {
	if id == "" {
		return model.Account{}, ErrInvalidInput
	}

	parsedID, err := uuid.Parse(string(id))
	if err != nil {
		return model.Account{}, ErrInvalidInput
	}

	var accountDTO dto.Account
	err = a.conn.WithContext(ctx).Where("id = ?", parsedID).First(&accountDTO).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Account{}, ErrRecordNotFound
		}
		return model.Account{}, err
	}

	return accountDTO.ToModel(), nil
}

func (a *accountPersistence) FindByUserID(ctx context.Context, userID string) (model.Account, error) {
	if userID == "" {
		return model.Account{}, ErrInvalidInput
	}

	var accountDTO dto.Account
	err := a.conn.WithContext(ctx).Where("user_uid = ?", userID).First(&accountDTO).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Account{}, ErrRecordNotFound
		}
		return model.Account{}, err
	}

	return accountDTO.ToModel(), nil
}

func (a *accountPersistence) Store(ctx context.Context, acc *model.Account) error {
	if acc == nil {
		return ErrInvalidInput
	}

	accountDTO := dto.ToAccountDto(*acc)

	// UPSERTの実装（存在すればUpdate、なければInsert）
	err := a.conn.WithContext(ctx).Save(&accountDTO).Error
	if err != nil {
		return err
	}

	return nil
}

func NewAccount(c *db.Conn) repository.Account {
	return &accountPersistence{conn: c}
}
