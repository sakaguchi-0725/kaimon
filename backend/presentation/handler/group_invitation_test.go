package handler_test

import (
	"backend/core"
	"backend/presentation/handler"
	mock "backend/test/mock/usecase"
	"backend/usecase"
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGroupInvitationHandler(t *testing.T) {

	tests := []struct {
		name         string
		groupID      string
		userID       string
		setupMock    func(*mock.MockGroupInvitation)
		expectedCode int
	}{
		{
			name:    "正常系：招待コードを生成",
			groupID: "01234567-89ab-cdef-0123-456789abcdef",
			userID:  "test-user-id",
			setupMock: func(m *mock.MockGroupInvitation) {
				input := usecase.GroupInvitationInput{
					GroupID: "01234567-89ab-cdef-0123-456789abcdef",
					UserID:  "test-user-id",
				}
				output := usecase.GroupInvitationOutput{
					Code:      "ABC123XY",
					ExpiresAt: "2024-01-08T12:30:45+09:00",
				}
				m.EXPECT().Execute(gomock.Any(), input).Return(output, nil)
			},
			expectedCode: http.StatusOK,
		},
		{
			name:    "異常系：存在しないユーザー",
			groupID: "01234567-89ab-cdef-0123-456789abcdef",
			userID:  "non-existent-user",
			setupMock: func(m *mock.MockGroupInvitation) {
				input := usecase.GroupInvitationInput{
					GroupID: "01234567-89ab-cdef-0123-456789abcdef",
					UserID:  "non-existent-user",
				}
				m.EXPECT().Execute(gomock.Any(), input).Return(usecase.GroupInvitationOutput{}, assert.AnError)
			},
			expectedCode: 0, // エラーが発生することを期待
		},
		{
			name:    "異常系：無効なグループID",
			groupID: "invalid-group-id",
			userID:  "test-user-id",
			setupMock: func(m *mock.MockGroupInvitation) {
				input := usecase.GroupInvitationInput{
					GroupID: "invalid-group-id",
					UserID:  "test-user-id",
				}
				m.EXPECT().Execute(gomock.Any(), input).Return(usecase.GroupInvitationOutput{}, assert.AnError)
			},
			expectedCode: 0, // エラーが発生することを期待
		},
		{
			name:    "異常系：管理者権限がない",
			groupID: "01234567-89ab-cdef-0123-456789abcdef",
			userID:  "test-user-id",
			setupMock: func(m *mock.MockGroupInvitation) {
				input := usecase.GroupInvitationInput{
					GroupID: "01234567-89ab-cdef-0123-456789abcdef",
					UserID:  "test-user-id",
				}
				forbiddenErr := core.NewAppError(core.ErrForbidden, assert.AnError).WithMessage("管理者権限がありません")
				m.EXPECT().Execute(gomock.Any(), input).Return(usecase.GroupInvitationOutput{}, forbiddenErr)
			},
			expectedCode: 0, // エラーが発生することを期待
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUsecase := mock.NewMockGroupInvitation(ctrl)
			tt.setupMock(mockUsecase)

			handlerFunc := handler.NewGroupInvitation(mockUsecase)

			// テストリクエストを作成
			rec, c := createTestRequest(http.MethodPost, "/groups/"+tt.groupID+"/invitations", "")
			
			// パラメータを設定
			c.SetPath("/groups/:id/invitations")
			c.SetParamNames("id")
			c.SetParamValues(tt.groupID)
			
			// ユーザーIDをコンテキストに設定
			ctx := context.WithValue(c.Request().Context(), core.UserIDKey, tt.userID)
			c.SetRequest(c.Request().WithContext(ctx))

			// ハンドラーを実行
			err := handlerFunc(c)

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

func TestGroupInvitationHandler_MissingUserID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGroupInvitation := mock.NewMockGroupInvitation(ctrl)

	// テストリクエストを作成（UserIDを設定しない）
	_, c := createTestRequest(http.MethodPost, "/groups/01234567-89ab-cdef-0123-456789abcdef/invitations", "")
	c.SetPath("/groups/:id/invitations")
	c.SetParamNames("id")
	c.SetParamValues("01234567-89ab-cdef-0123-456789abcdef")

	// ハンドラーを取得して実行
	handlerFunc := handler.NewGroupInvitation(mockGroupInvitation)
	
	// UserIDが設定されていない場合、core.GetUserIDが空文字を返すので、
	// そのままusecaseに渡される
	input := usecase.GroupInvitationInput{
		GroupID: "01234567-89ab-cdef-0123-456789abcdef",
		UserID:  "", // 空文字
	}
	mockGroupInvitation.EXPECT().Execute(gomock.Any(), input).Return(usecase.GroupInvitationOutput{}, assert.AnError)

	err := handlerFunc(c)
	
	// エラーが返されることを確認
	assert.Error(t, err)
}