package model_test

import (
	"backend/domain/model"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAccount(t *testing.T) {
	t.Run("NewAccount", func(t *testing.T) {
		validID := model.AccountID(uuid.NewString())
		validUserID := "user123"
		validName := "テストユーザー"

		tests := []struct {
			name    string
			id      model.AccountID
			userID  string
			accName string
			want    model.Account
			wantErr error
		}{
			{
				name:    "正常なアカウント作成",
				id:      validID,
				userID:  validUserID,
				accName: validName,
				want: model.Account{
					ID:     validID,
					UserID: validUserID,
					Name:   validName,
				},
				wantErr: nil,
			},
			{
				name:    "userIDが空文字の場合",
				id:      validID,
				userID:  "",
				accName: validName,
				want:    model.Account{},
				wantErr: errors.New("userID is required"),
			},
			{
				name:    "nameが空文字の場合",
				id:      validID,
				userID:  validUserID,
				accName: "",
				want:    model.Account{},
				wantErr: errors.New("name is required"),
			},
			{
				name:    "userIDとnameが両方空文字の場合",
				id:      validID,
				userID:  "",
				accName: "",
				want:    model.Account{},
				wantErr: errors.New("userID is required"),
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := model.NewAccount(tt.id, tt.userID, tt.accName)

				assert.Equal(t, tt.want, got)
				if tt.wantErr != nil {
					assert.Equal(t, tt.wantErr.Error(), err.Error())
				} else {
					assert.NoError(t, err)
				}
			})
		}
	})
}