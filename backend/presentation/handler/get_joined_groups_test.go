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

func TestGetJoinedGroups(t *testing.T) {
	tests := []struct {
		name         string
		userID       string
		setupMock    func(*mock.MockGetJoinedGroups)
		expectedCode int
		expectedBody string
	}{
		{
			name:   "正常ケース: ユーザーが複数のグループに参加している場合",
			userID: "test-user-id",
			setupMock: func(m *mock.MockGetJoinedGroups) {
				m.EXPECT().Execute(gomock.Any(), "test-user-id").Return([]usecase.GetJoinedGroupOutput{
					{
						ID:          "group-1",
						Name:        "テストグループ1",
						MemberCount: 5,
					},
					{
						ID:          "group-2",
						Name:        "テストグループ2",
						MemberCount: 3,
					},
				}, nil)
			},
			expectedCode: http.StatusOK,
			expectedBody: `{"groups":[{"id":"group-1","name":"テストグループ1","memberCount":5},{"id":"group-2","name":"テストグループ2","memberCount":3}]}`,
		},
		{
			name:   "正常ケース: ユーザーがグループに参加していない場合",
			userID: "test-user-id",
			setupMock: func(m *mock.MockGetJoinedGroups) {
				m.EXPECT().Execute(gomock.Any(), "test-user-id").Return([]usecase.GetJoinedGroupOutput{}, nil)
			},
			expectedCode: http.StatusOK,
			expectedBody: `{"groups":[]}`,
		},
		{
			name:   "正常ケース: 1つのグループに参加している場合",
			userID: "test-user-id",
			setupMock: func(m *mock.MockGetJoinedGroups) {
				m.EXPECT().Execute(gomock.Any(), "test-user-id").Return([]usecase.GetJoinedGroupOutput{
					{
						ID:          "group-1",
						Name:        "ソログループ",
						MemberCount: 1,
					},
				}, nil)
			},
			expectedCode: http.StatusOK,
			expectedBody: `{"groups":[{"id":"group-1","name":"ソログループ","memberCount":1}]}`,
		},
		{
			name:         "異常ケース: ユーザーIDがコンテキストにない場合",
			userID:       "",
			setupMock:    func(m *mock.MockGetJoinedGroups) {},
			expectedCode: http.StatusUnauthorized,
		},
		{
			name:   "異常ケース: UseCase実行エラー",
			userID: "test-user-id",
			setupMock: func(m *mock.MockGetJoinedGroups) {
				m.EXPECT().Execute(gomock.Any(), "test-user-id").Return(nil, errors.New("usecase error"))
			},
			expectedCode: http.StatusInternalServerError,
		},
		{
			name:   "異常ケース: アカウントが見つからない場合",
			userID: "non-existent-user",
			setupMock: func(m *mock.MockGetJoinedGroups) {
				m.EXPECT().Execute(gomock.Any(), "non-existent-user").Return(nil, errors.New("account not found"))
			},
			expectedCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUseCase := mock.NewMockGetJoinedGroups(ctrl)
			tt.setupMock(mockUseCase)

			handler := handler.NewGetJoinedGroups(mockUseCase)

			rec, c := createTestGetRequest("/groups/joined")

			// コンテキストにユーザー情報を設定
			ctx := c.Request().Context()
			if tt.userID != "" {
				ctx = context.WithValue(ctx, core.UserIDKey, tt.userID)
			}
			c.SetRequest(c.Request().WithContext(ctx))

			err := handler(c)

			if tt.expectedCode == http.StatusOK {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedCode, rec.Code)
				assert.JSONEq(t, tt.expectedBody, rec.Body.String())
			} else {
				assert.Error(t, err)
			}
		})
	}
}