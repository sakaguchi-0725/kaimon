package handler_test

import (
	"backend/core"
	"backend/presentation/handler"
	mock "backend/test/mock/usecase"
	"backend/usecase"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name           string
		userID         string
		setupMock      func(*mock.MockGetAccount)
		expectedStatus int
		expectedBody   string
	}{
		{
			name:   "正常系：アカウント情報を取得",
			userID: "test-user-id",
			setupMock: func(m *mock.MockGetAccount) {
				m.EXPECT().Execute(gomock.Any(), usecase.GetAccountInput{
					UserID: "test-user-id",
				}).Return(usecase.GetAccountOutput{
					ID:   "test-account-id",
					Name: "テストユーザー",
				}, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: `{
				"id": "test-account-id",
				"name": "テストユーザー"
			}`,
		},
		{
			name:   "正常系：長い名前のアカウント",
			userID: "long-name-user",
			setupMock: func(m *mock.MockGetAccount) {
				m.EXPECT().Execute(gomock.Any(), usecase.GetAccountInput{
					UserID: "long-name-user",
				}).Return(usecase.GetAccountOutput{
					ID:   "long-name-account-id",
					Name: "これは非常に長い名前のテストユーザーです",
				}, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: `{
				"id": "long-name-account-id",
				"name": "これは非常に長い名前のテストユーザーです"
			}`,
		},
		{
			name:   "異常系：UseCaseでエラーが発生（アカウントが見つからない）",
			userID: "non-existent-user",
			setupMock: func(m *mock.MockGetAccount) {
				m.EXPECT().Execute(gomock.Any(), usecase.GetAccountInput{
					UserID: "non-existent-user",
				}).Return(usecase.GetAccountOutput{}, errors.New("account not found"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "",
		},
		{
			name:   "異常系：UseCaseでデータベースエラーが発生",
			userID: "db-error-user",
			setupMock: func(m *mock.MockGetAccount) {
				m.EXPECT().Execute(gomock.Any(), usecase.GetAccountInput{
					UserID: "db-error-user",
				}).Return(usecase.GetAccountOutput{}, errors.New("database error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			mockUseCase := mock.NewMockGetAccount(ctrl)
			tt.setupMock(mockUseCase)

			rec, c := createTestGetRequest("/account")
			
			// ユーザーIDをコンテキストに設定
			ctx := context.WithValue(c.Request().Context(), core.UserIDKey, tt.userID)
			c.SetRequest(c.Request().WithContext(ctx))

			handlerFunc := handler.NewGetAccount(mockUseCase)

			// Execute
			err := handlerFunc(c)

			// Assert
			if tt.expectedStatus >= 400 {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedStatus, rec.Code)

				if tt.expectedBody != "" {
					var expected, actual map[string]interface{}
					err := json.Unmarshal([]byte(tt.expectedBody), &expected)
					assert.NoError(t, err)
					err = json.Unmarshal(rec.Body.Bytes(), &actual)
					assert.NoError(t, err)
					assert.Equal(t, expected, actual)
				}
			}
		})
	}
}