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

func TestGetGroupMembers(t *testing.T) {
	tests := []struct {
		name         string
		groupID      string
		userID       string
		setupMock    func(*mock.MockGetGroupMembers)
		expectedCode int
		expectedBody string
	}{
		{
			name:    "正常ケース: 複数メンバー取得",
			groupID: "group-123",
			userID:  "user-123",
			setupMock: func(m *mock.MockGetGroupMembers) {
				m.EXPECT().Execute(gomock.Any(), usecase.GetGroupMembersInput{
					GroupID: "group-123",
					UserID:  "user-123",
				}).Return(usecase.GetGroupMembersOutput{
					Members: []usecase.Member{
						{
							ID:     "member-1",
							Name:   "管理者",
							Role:   "admin",
							Status: "active",
						},
						{
							ID:     "member-2",
							Name:   "メンバー1",
							Role:   "member",
							Status: "active",
						},
						{
							ID:     "member-3",
							Name:   "メンバー2",
							Role:   "member",
							Status: "pending",
						},
					},
				}, nil)
			},
			expectedCode: http.StatusOK,
			expectedBody: `{"members":[{"id":"member-1","name":"管理者","role":"admin","status":"active"},{"id":"member-2","name":"メンバー1","role":"member","status":"active"},{"id":"member-3","name":"メンバー2","role":"member","status":"pending"}]}`,
		},
		{
			name:    "正常ケース: 単一メンバー",
			groupID: "group-456",
			userID:  "user-456",
			setupMock: func(m *mock.MockGetGroupMembers) {
				m.EXPECT().Execute(gomock.Any(), usecase.GetGroupMembersInput{
					GroupID: "group-456",
					UserID:  "user-456",
				}).Return(usecase.GetGroupMembersOutput{
					Members: []usecase.Member{
						{
							ID:     "member-1",
							Name:   "作成者",
							Role:   "admin",
							Status: "active",
						},
					},
				}, nil)
			},
			expectedCode: http.StatusOK,
			expectedBody: `{"members":[{"id":"member-1","name":"作成者","role":"admin","status":"active"}]}`,
		},
		{
			name:    "異常ケース: GroupID不正",
			groupID: "invalid-group-id",
			userID:  "user-123",
			setupMock: func(m *mock.MockGetGroupMembers) {
				m.EXPECT().Execute(gomock.Any(), usecase.GetGroupMembersInput{
					GroupID: "invalid-group-id",
					UserID:  "user-123",
				}).Return(usecase.GetGroupMembersOutput{}, errors.New("invalid group id"))
			},
			expectedCode: http.StatusInternalServerError,
		},
		{
			name:    "異常ケース: 権限なし",
			groupID: "group-123",
			userID:  "unauthorized-user",
			setupMock: func(m *mock.MockGetGroupMembers) {
				m.EXPECT().Execute(gomock.Any(), usecase.GetGroupMembersInput{
					GroupID: "group-123",
					UserID:  "unauthorized-user",
				}).Return(usecase.GetGroupMembersOutput{}, errors.New("not a member of the group"))
			},
			expectedCode: http.StatusInternalServerError,
		},
		{
			name:    "異常ケース: グループ不存在",
			groupID: "non-existent-group",
			userID:  "user-123",
			setupMock: func(m *mock.MockGetGroupMembers) {
				m.EXPECT().Execute(gomock.Any(), usecase.GetGroupMembersInput{
					GroupID: "non-existent-group",
					UserID:  "user-123",
				}).Return(usecase.GetGroupMembersOutput{}, errors.New("record not found"))
			},
			expectedCode: http.StatusInternalServerError,
		},
		{
			name:    "異常ケース: UseCase実行エラー",
			groupID: "group-123",
			userID:  "user-123",
			setupMock: func(m *mock.MockGetGroupMembers) {
				m.EXPECT().Execute(gomock.Any(), usecase.GetGroupMembersInput{
					GroupID: "group-123",
					UserID:  "user-123",
				}).Return(usecase.GetGroupMembersOutput{}, errors.New("internal server error"))
			},
			expectedCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUseCase := mock.NewMockGetGroupMembers(ctrl)
			tt.setupMock(mockUseCase)

			handler := handler.NewGetGroupMembers(mockUseCase)

			rec, c := createTestGetRequest("/groups/" + tt.groupID + "/members")
			c.SetParamNames("id")
			c.SetParamValues(tt.groupID)

			// コンテキストにユーザー情報を設定
			ctx := context.WithValue(c.Request().Context(), core.UserIDKey, tt.userID)
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