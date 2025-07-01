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

func TestGetGroup(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name           string
		groupID        string
		userID         string
		setupMock      func(*mock.MockGetGroup)
		expectedStatus int
		expectedBody   string
	}{
		{
			name:    "正常系：グループ詳細とメンバー一覧を取得",
			groupID: "test-group-id",
			userID:  "test-user-id",
			setupMock: func(m *mock.MockGetGroup) {
				m.EXPECT().Execute(gomock.Any(), usecase.GetGroupInput{
					UserID:  "test-user-id",
					GroupID: "test-group-id",
				}).Return(usecase.GetGroupOutput{
					ID:          "test-group-id",
					Name:        "テストグループ",
					Description: "テストグループの説明",
					CreatedAt:   "2023-01-01T00:00:00Z",
					Members: []usecase.Member{
						{
							ID:     "member1",
							Name:   "メンバー1",
							Role:   "admin",
							Status: "active",
						},
						{
							ID:     "member2",
							Name:   "メンバー2",
							Role:   "member",
							Status: "active",
						},
					},
				}, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: `{
				"id": "test-group-id",
				"name": "テストグループ",
				"description": "テストグループの説明",
				"createdAt": "2023-01-01T00:00:00Z",
				"members": [
					{
						"id": "member1",
						"name": "メンバー1",
						"role": "admin",
						"status": "active"
					},
					{
						"id": "member2",
						"name": "メンバー2",
						"role": "member",
						"status": "active"
					}
				]
			}`,
		},
		{
			name:    "正常系：メンバーが1人のグループ",
			groupID: "single-member-group",
			userID:  "test-user-id",
			setupMock: func(m *mock.MockGetGroup) {
				m.EXPECT().Execute(gomock.Any(), usecase.GetGroupInput{
					UserID:  "test-user-id",
					GroupID: "single-member-group",
				}).Return(usecase.GetGroupOutput{
					ID:          "single-member-group",
					Name:        "単一メンバーグループ",
					Description: "作成者のみのグループ",
					CreatedAt:   "2023-01-01T00:00:00Z",
					Members: []usecase.Member{
						{
							ID:     "member1",
							Name:   "作成者",
							Role:   "admin",
							Status: "active",
						},
					},
				}, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: `{
				"id": "single-member-group",
				"name": "単一メンバーグループ",
				"description": "作成者のみのグループ",
				"createdAt": "2023-01-01T00:00:00Z",
				"members": [
					{
						"id": "member1",
						"name": "作成者",
						"role": "admin",
						"status": "active"
					}
				]
			}`,
		},
		{
			name:    "異常系：UseCaseでエラーが発生",
			groupID: "error-group",
			userID:  "test-user-id",
			setupMock: func(m *mock.MockGetGroup) {
				m.EXPECT().Execute(gomock.Any(), usecase.GetGroupInput{
					UserID:  "test-user-id",
					GroupID: "error-group",
				}).Return(usecase.GetGroupOutput{}, errors.New("グループが見つかりません"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			mockUseCase := mock.NewMockGetGroup(ctrl)
			tt.setupMock(mockUseCase)

			rec, c := createTestGetRequest("/groups/" + tt.groupID)
			
			// ユーザーIDをコンテキストに設定
			ctx := context.WithValue(c.Request().Context(), core.UserIDKey, tt.userID)
			c.SetRequest(c.Request().WithContext(ctx))
			
			// パラメータを設定
			c.SetParamNames("id")
			c.SetParamValues(tt.groupID)

			handlerFunc := handler.NewGetGroup(mockUseCase)

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