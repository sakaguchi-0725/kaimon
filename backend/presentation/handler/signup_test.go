package handler_test

import (
	"backend/core"
	"backend/presentation/handler"
	mock "backend/test/mock/usecase"
	"backend/usecase"
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestSignUp(t *testing.T) {
	tests := []struct {
		name         string
		requestBody  string
		userID       string
		email        string
		setupMock    func(*mock.MockSignUp)
		expectedCode int
	}{
		{
			name:        "正常なリクエスト",
			requestBody: `{"name":"テストユーザー","profileImageUrl":"https://example.com/image.jpg"}`,
			userID:      "firebase-uid-123",
			email:       "test@example.com",
			setupMock: func(m *mock.MockSignUp) {
				m.EXPECT().Execute(gomock.Any(), usecase.SignUpInput{
					UID:             "firebase-uid-123",
					Email:           "test@example.com", 
					Name:            "テストユーザー",
					ProfileImageURL: "https://example.com/image.jpg",
				}).Return(nil)
			},
			expectedCode: http.StatusNoContent,
		},
		{
			name:        "プロフィール画像URLなし",
			requestBody: `{"name":"テストユーザー"}`,
			userID:      "firebase-uid-123",
			email:       "test@example.com",
			setupMock: func(m *mock.MockSignUp) {
				m.EXPECT().Execute(gomock.Any(), usecase.SignUpInput{
					UID:             "firebase-uid-123",
					Email:           "test@example.com",
					Name:            "テストユーザー",
					ProfileImageURL: "",
				}).Return(nil)
			},
			expectedCode: http.StatusNoContent,
		},
		{
			name:         "ユーザーIDがコンテキストにない",
			requestBody:  `{"name":"テストユーザー"}`,
			userID:       "",
			email:        "test@example.com",
			setupMock:    func(m *mock.MockSignUp) {},
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "emailがコンテキストにない",
			requestBody:  `{"name":"テストユーザー"}`,
			userID:       "firebase-uid-123",
			email:        "",
			setupMock:    func(m *mock.MockSignUp) {},
			expectedCode: http.StatusBadRequest,
		},
		{
			name:        "不正なJSONリクエスト",
			requestBody: `{"name":}`,
			userID:      "firebase-uid-123",
			email:       "test@example.com",
			setupMock:   func(m *mock.MockSignUp) {},
			expectedCode: http.StatusBadRequest,
		},
		{
			name:        "UseCase実行エラー",
			requestBody: `{"name":"テストユーザー"}`,
			userID:      "firebase-uid-123",
			email:       "test@example.com",
			setupMock: func(m *mock.MockSignUp) {
				m.EXPECT().Execute(gomock.Any(), usecase.SignUpInput{
					UID:             "firebase-uid-123",
					Email:           "test@example.com",
					Name:            "テストユーザー",
					ProfileImageURL: "",
				}).Return(errors.New("usecase error"))
			},
			expectedCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUseCase := mock.NewMockSignUp(ctrl)
			tt.setupMock(mockUseCase)

			handler := handler.NewSignUp(mockUseCase)

			rec, c := createTestPostRequest("/signup", tt.requestBody)
			
			// コンテキストにユーザー情報を設定
			ctx := c.Request().Context()
			if tt.userID != "" {
				ctx = context.WithValue(ctx, core.UserIDKey, tt.userID)
			}
			if tt.email != "" {
				ctx = context.WithValue(ctx, core.EmailKey, tt.email)
			}
			c.SetRequest(c.Request().WithContext(ctx))

			err := handler(c)

			if tt.expectedCode == http.StatusNoContent {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedCode, rec.Code)
			} else {
				assert.Error(t, err)
			}
		})
	}
}