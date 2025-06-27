package usecase_test

import (
	"backend/infra/persistence"
	mock "backend/test/mock/repository"
	"backend/usecase"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestVerifyToken_Execute(t *testing.T) {
	tests := []struct {
		name      string
		token     string
		setupMock func(*mock.MockAuthenticator)
		wantUID   string
		wantEmail string
		wantErr   error
	}{
		{
			name:  "正常なトークン検証",
			token: "valid_token",
			setupMock: func(auth *mock.MockAuthenticator) {
				auth.EXPECT().VerifyToken("valid_token").Return("firebase-uid-123", "test@example.com", nil)
			},
			wantUID:   "firebase-uid-123",
			wantEmail: "test@example.com",
			wantErr:   nil,
		},
		{
			name:  "無効なトークン",
			token: "invalid_token",
			setupMock: func(auth *mock.MockAuthenticator) {
				auth.EXPECT().VerifyToken("invalid_token").Return("", "", persistence.ErrInvalidToken)
			},
			wantUID:   "",
			wantEmail: "",
			wantErr:   persistence.ErrInvalidToken,
		},
		{
			name:  "期限切れトークン",
			token: "expired_token",
			setupMock: func(auth *mock.MockAuthenticator) {
				auth.EXPECT().VerifyToken("expired_token").Return("", "", errors.New("token expired"))
			},
			wantUID:   "",
			wantEmail: "",
			wantErr:   errors.New("token expired"),
		},
		{
			name:  "空のトークン",
			token: "",
			setupMock: func(auth *mock.MockAuthenticator) {
				auth.EXPECT().VerifyToken("").Return("", "", errors.New("token is required"))
			},
			wantUID:   "",
			wantEmail: "",
			wantErr:   errors.New("token is required"),
		},
		{
			name:  "認証サービス内部エラー",
			token: "service_error_token",
			setupMock: func(auth *mock.MockAuthenticator) {
				auth.EXPECT().VerifyToken("service_error_token").Return("", "", errors.New("internal authentication error"))
			},
			wantUID:   "",
			wantEmail: "",
			wantErr:   errors.New("internal authentication error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockAuth := mock.NewMockAuthenticator(ctrl)
			verifyToken := usecase.NewVerifyToken(mockAuth)

			tt.setupMock(mockAuth)

			uid, email, err := verifyToken.Execute(context.Background(), tt.token)

			if tt.wantErr != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.wantErr.Error())
				assert.Equal(t, tt.wantUID, uid)
				assert.Equal(t, tt.wantEmail, email)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantUID, uid)
				assert.Equal(t, tt.wantEmail, email)
			}
		})
	}
}
