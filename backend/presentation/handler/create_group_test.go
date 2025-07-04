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

func TestCreateGroup(t *testing.T) {
	tests := []struct {
		name         string
		requestBody  string
		userID       string
		setupMock    func(*mock.MockCreateGroup)
		expectedCode int
	}{
		{
			name:        "正常なリクエスト",
			requestBody: `{"name":"テストグループ","description":"テスト用のグループです"}`,
			userID:      "firebase-uid-123",
			setupMock: func(m *mock.MockCreateGroup) {
				m.EXPECT().Execute(gomock.Any(), usecase.CreateGroupInput{
					UserID:      "firebase-uid-123",
					Name:        "テストグループ",
					Description: "テスト用のグループです",
				}).Return(nil)
			},
			expectedCode: http.StatusCreated,
		},
		{
			name:        "説明なしのリクエスト",
			requestBody: `{"name":"説明なしグループ"}`,
			userID:      "firebase-uid-123",
			setupMock: func(m *mock.MockCreateGroup) {
				m.EXPECT().Execute(gomock.Any(), usecase.CreateGroupInput{
					UserID:      "firebase-uid-123",
					Name:        "説明なしグループ",
					Description: "",
				}).Return(nil)
			},
			expectedCode: http.StatusCreated,
		},
		{
			name:        "空の説明フィールド",
			requestBody: `{"name":"空説明グループ","description":""}`,
			userID:      "firebase-uid-123",
			setupMock: func(m *mock.MockCreateGroup) {
				m.EXPECT().Execute(gomock.Any(), usecase.CreateGroupInput{
					UserID:      "firebase-uid-123",
					Name:        "空説明グループ",
					Description: "",
				}).Return(nil)
			},
			expectedCode: http.StatusCreated,
		},
		{
			name:        "ユースケースエラー",
			requestBody: `{"name":"エラーグループ","description":"エラーテスト"}`,
			userID:      "firebase-uid-123",
			setupMock: func(m *mock.MockCreateGroup) {
				m.EXPECT().Execute(gomock.Any(), usecase.CreateGroupInput{
					UserID:      "firebase-uid-123",
					Name:        "エラーグループ",
					Description: "エラーテスト",
				}).Return(errors.New("usecase error"))
			},
			expectedCode: 0, // エラーが発生することを期待
		},
		{
			name:         "不正なJSON",
			requestBody:  `{"name":}`,
			userID:       "firebase-uid-123",
			setupMock:    func(m *mock.MockCreateGroup) {},
			expectedCode: 0, // エラーが発生することを期待
		},
		{
			name:         "空のリクエストボディ",
			requestBody:  `{}`,
			userID:       "firebase-uid-123",
			setupMock: func(m *mock.MockCreateGroup) {
				m.EXPECT().Execute(gomock.Any(), usecase.CreateGroupInput{
					UserID:      "firebase-uid-123",
					Name:        "",
					Description: "",
				}).Return(core.NewInvalidError(errors.New("name is required")))
			},
			expectedCode: 0, // エラーが発生することを期待
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUsecase := mock.NewMockCreateGroup(ctrl)
			tt.setupMock(mockUsecase)

			handler := handler.NewCreateGroup(mockUsecase)

			rec, c := createTestRequest(http.MethodPost, "/groups", tt.requestBody)
			ctx := context.WithValue(c.Request().Context(), core.UserIDKey, tt.userID)
			c.SetRequest(c.Request().WithContext(ctx))

			err := handler(c)
			
			// エラーハンドリングはmiddlewareで処理されるため、
			// ここではエラーの有無のみ確認
			if tt.expectedCode == 0 {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedCode, rec.Code)
			}
		})
	}
}