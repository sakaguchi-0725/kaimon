package persistence_test

import (
	"backend/infra/persistence"
	mock "backend/test/mock/external"
	"context"
	"errors"
	"testing"

	"firebase.google.com/go/v4/auth"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestAuthenticator_VerifyToken(t *testing.T) {
	tests := []struct {
		name      string
		token     string
		setupMock func(*mock.MockFirebaseClient)
		wantUID   string
		wantEmail string
		wantErr   bool
	}{
		{
			name:  "正常なトークン（メール/パスワード認証）",
			token: "valid_email_password_token",
			setupMock: func(m *mock.MockFirebaseClient) {
				authToken := &auth.Token{
					UID: "test-uid-123",
					Claims: map[string]interface{}{
						"email":          "user@example.com",
						"email_verified": false,
						"displayName":    "テストユーザー",
					},
				}
				m.EXPECT().VerifyIDToken(gomock.Any(), "valid_email_password_token").Return(authToken, nil)
			},
			wantUID:   "test-uid-123",
			wantEmail: "user@example.com",
			wantErr:   false,
		},
		{
			name:  "正常なトークン（Google認証）",
			token: "valid_google_token",
			setupMock: func(m *mock.MockFirebaseClient) {
				authToken := &auth.Token{
					UID: "test-uid-456",
					Claims: map[string]interface{}{
						"email":          "googleuser@gmail.com",
						"email_verified": true,
						"displayName":    "Google User",
					},
				}
				m.EXPECT().VerifyIDToken(gomock.Any(), "valid_google_token").Return(authToken, nil)
			},
			wantUID:   "test-uid-456",
			wantEmail: "googleuser@gmail.com",
			wantErr:   false,
		},
		{
			name:  "無効なトークンの場合",
			token: "invalid_token",
			setupMock: func(m *mock.MockFirebaseClient) {
				m.EXPECT().VerifyIDToken(gomock.Any(), "invalid_token").Return(nil, errors.New("invalid token"))
			},
			wantUID:   "",
			wantEmail: "",
			wantErr:   true,
		},
		{
			name:  "空文字のトークンの場合",
			token: "",
			setupMock: func(m *mock.MockFirebaseClient) {
				m.EXPECT().VerifyIDToken(gomock.Any(), "").Return(nil, errors.New("empty token"))
			},
			wantUID:   "",
			wantEmail: "",
			wantErr:   true,
		},
		{
			name:  "Firebase接続エラーの場合",
			token: "connection_error_token",
			setupMock: func(m *mock.MockFirebaseClient) {
				m.EXPECT().VerifyIDToken(gomock.Any(), "connection_error_token").Return(nil, errors.New("firebase connection error"))
			},
			wantUID:   "",
			wantEmail: "",
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockClient := mock.NewMockFirebaseClient(ctrl)
			tt.setupMock(mockClient)

			authenticator := persistence.NewAuthenticator(mockClient)
			gotUID, gotEmail, err := authenticator.VerifyToken(context.Background(), tt.token)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Empty(t, gotUID)
				assert.Empty(t, gotEmail)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantUID, gotUID)
				assert.Equal(t, tt.wantEmail, gotEmail)
			}
		})
	}
}
