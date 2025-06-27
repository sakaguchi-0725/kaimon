package model_test

import (
	"backend/core"
	"backend/domain/model"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	t.Run("NewUser", func(t *testing.T) {
		validID := "user123"
		validEmail := "test@example.com"

		tests := []struct {
			name    string
			id      string
			email   string
			want    model.User
			wantErr *core.AppError
		}{
			{
				name:  "正常なユーザー作成",
				id:    validID,
				email: validEmail,
				want: model.User{
					ID:    validID,
					Email: validEmail,
				},
				wantErr: nil,
			},
			{
				name:    "idが空文字の場合",
				id:      "",
				email:   validEmail,
				want:    model.User{},
				wantErr: core.NewInvalidError(errors.New("id is required")),
			},
			{
				name:    "emailが空文字の場合",
				id:      validID,
				email:   "",
				want:    model.User{},
				wantErr: core.NewInvalidError(errors.New("email is required")),
			},
			{
				name:    "idとemailが両方空文字の場合",
				id:      "",
				email:   "",
				want:    model.User{},
				wantErr: core.NewInvalidError(errors.New("id is required")),
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := model.NewUser(tt.id, tt.email)

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