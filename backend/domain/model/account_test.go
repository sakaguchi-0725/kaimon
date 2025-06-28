package model_test

import (
	"backend/core"
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
			name            string
			id              model.AccountID
			userID          string
			accName         string
			profileImageURL string
			want            model.Account
			wantErr         *core.AppError
		}{
			{
				name:            "正常なアカウント作成",
				id:              validID,
				userID:          validUserID,
				accName:         validName,
				profileImageURL: "https://example.com/image.jpg",
				want: model.Account{
					ID:              validID,
					UserID:          validUserID,
					Name:            validName,
					ProfileImageURL: "https://example.com/image.jpg",
				},
				wantErr: nil,
			},
			{
				name:            "userIDが空文字の場合",
				id:              validID,
				userID:          "",
				accName:         validName,
				profileImageURL: "https://example.com/image.jpg",
				want:            model.Account{},
				wantErr:         core.NewInvalidError(errors.New("userID is required")),
			},
			{
				name:            "nameが空文字の場合",
				id:              validID,
				userID:          validUserID,
				accName:         "",
				profileImageURL: "https://example.com/image.jpg",
				want:            model.Account{},
				wantErr:         core.NewInvalidError(errors.New("name is required")),
			},
			{
				name:            "userIDとnameが両方空文字の場合",
				id:              validID,
				userID:          "",
				accName:         "",
				profileImageURL: "https://example.com/image.jpg",
				want:            model.Account{},
				wantErr:         core.NewInvalidError(errors.New("userID is required")),
			},
			{
				name:            "profileImageURLが空の場合",
				id:              validID,
				userID:          validUserID,
				accName:         validName,
				profileImageURL: "",
				want: model.Account{
					ID:              validID,
					UserID:          validUserID,
					Name:            validName,
					ProfileImageURL: "",
				},
				wantErr: nil,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := model.NewAccount(tt.id, tt.userID, tt.accName, tt.profileImageURL)

				assert.Equal(t, tt.want, got)
				if tt.wantErr != nil {
					assert.Error(t, err)
					assert.ErrorIs(t, err, tt.wantErr)
					// エラーコードも検証
					var appErr *core.AppError
					if assert.ErrorAs(t, err, &appErr) {
						assert.Equal(t, tt.wantErr.Code(), appErr.Code())
						assert.Equal(t, tt.wantErr.Unwrap().Error(), appErr.Unwrap().Error())
					}
				} else {
					assert.NoError(t, err)
				}
			})
		}
	})
}