package usecase_test

import (
	"backend/domain/model"
	"backend/infra/persistence"
	mock "backend/test/mock/repository"
	"backend/usecase"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestLogin_Execute(t *testing.T) {
	tests := []struct {
		name      string
		input     usecase.LoginInput
		setupMock func(*mock.MockAuthenticator, *mock.MockUser)
		wantErr   error
	}{
		{
			name: "正常なログイン",
			input: usecase.LoginInput{
				Email:    "test@example.com",
				Password: "password123",
			},
			setupMock: func(auth *mock.MockAuthenticator, user *mock.MockUser) {
				ctx := context.Background()
				
				// Firebase Auth認証成功
				auth.EXPECT().SignInWithEmailAndPassword("test@example.com", "password123").Return("firebase-uid-123", nil)
				
				// ユーザー存在確認成功
				expectedUser, _ := model.NewUser("firebase-uid-123", "test@example.com")
				user.EXPECT().FindByUID(ctx, "firebase-uid-123").Return(&expectedUser, nil)
			},
			wantErr: nil,
		},
		{
			name: "認証失敗（無効なメール/パスワード）",
			input: usecase.LoginInput{
				Email:    "test@example.com",
				Password: "wrongpassword",
			},
			setupMock: func(auth *mock.MockAuthenticator, user *mock.MockUser) {
				// Firebase Auth認証失敗
				auth.EXPECT().SignInWithEmailAndPassword("test@example.com", "wrongpassword").Return("", persistence.ErrInvalidCredentials)
			},
			wantErr: persistence.ErrInvalidCredentials,
		},
		{
			name: "ユーザーが存在しない",
			input: usecase.LoginInput{
				Email:    "test@example.com",
				Password: "password123",
			},
			setupMock: func(auth *mock.MockAuthenticator, user *mock.MockUser) {
				ctx := context.Background()
				
				// Firebase Auth認証成功
				auth.EXPECT().SignInWithEmailAndPassword("test@example.com", "password123").Return("firebase-uid-123", nil)
				
				// ユーザーが見つからない
				user.EXPECT().FindByUID(ctx, "firebase-uid-123").Return(nil, persistence.ErrNotFound)
			},
			wantErr: persistence.ErrNotFound,
		},
		{
			name: "Authenticatorエラー（システムエラー）",
			input: usecase.LoginInput{
				Email:    "test@example.com",
				Password: "password123",
			},
			setupMock: func(auth *mock.MockAuthenticator, user *mock.MockUser) {
				// Firebase Auth システムエラー
				auth.EXPECT().SignInWithEmailAndPassword("test@example.com", "password123").Return("", errors.New("firebase auth system error"))
			},
			wantErr: errors.New("firebase auth system error"),
		},
		{
			name: "ユーザー検索エラー（システムエラー）",
			input: usecase.LoginInput{
				Email:    "test@example.com",
				Password: "password123",
			},
			setupMock: func(auth *mock.MockAuthenticator, user *mock.MockUser) {
				ctx := context.Background()
				
				// Firebase Auth認証成功
				auth.EXPECT().SignInWithEmailAndPassword("test@example.com", "password123").Return("firebase-uid-123", nil)
				
				// ユーザー検索でシステムエラー
				user.EXPECT().FindByUID(ctx, "firebase-uid-123").Return(nil, errors.New("database connection error"))
			},
			wantErr: errors.New("database connection error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockAuth := mock.NewMockAuthenticator(ctrl)
			mockUser := mock.NewMockUser(ctrl)

			login := usecase.NewLogin(mockAuth, mockUser)

			tt.setupMock(mockAuth, mockUser)

			err := login.Execute(context.Background(), tt.input)

			if tt.wantErr != nil {
				assert.Error(t, err)
				// エラーメッセージの部分一致で確認
				assert.Contains(t, err.Error(), tt.wantErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

