package model_test

import (
	"backend/domain/model"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAccountID(t *testing.T) {
	t.Run("ParseAccountID", func(t *testing.T) {
		validID := uuid.NewString()
		invalidID := "invalid-id"

		tests := []struct {
			name    string
			input   string
			want    model.AccountID
			wantErr error
		}{
			{
				name:    "正常なID",
				input:   validID,
				want:    model.AccountID(validID),
				wantErr: nil,
			},
			{
				name:    "不正なID",
				input:   invalidID,
				want:    "",
				wantErr: errors.New("invalid UUID"),
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := model.ParseAccountID(tt.input)

				assert.Equal(t, tt.want, got)
				if tt.wantErr != nil {
					assert.Error(t, err)
				} else {
					assert.NoError(t, err)
				}
			})
		}
	})
}
