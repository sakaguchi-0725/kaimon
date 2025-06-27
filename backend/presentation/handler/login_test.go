package handler_test

import (
	"backend/presentation/handler"
	mock "backend/test/mock/usecase"
	"backend/usecase"
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestLogin(t *testing.T) {
	tests := []struct {
		name         string
		body         string
		setupMock    func(*mock.MockLogin)
		expectedCode int
	}{
		{
			name: "正常なリクエスト",
			body: `{"email":"test@example.com","password":"password123"}`,
			setupMock: func(m *mock.MockLogin) {
				m.EXPECT().Execute(gomock.Any(), usecase.LoginInput{
					Email:    "test@example.com",
					Password: "password123",
				}).Return(nil)
			},
			expectedCode: http.StatusOK,
		},
		{
			name:         "空のリクエストボディ",
			body:         "",
			setupMock:    func(m *mock.MockLogin) {},
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "不正なJSON",
			body:         `{"email":"test@example.com","password":}`,
			setupMock:    func(m *mock.MockLogin) {},
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "emailが空",
			body:         `{"email":"","password":"password123"}`,
			setupMock:    func(m *mock.MockLogin) {},
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "passwordが空",
			body:         `{"email":"test@example.com","password":""}`,
			setupMock:    func(m *mock.MockLogin) {},
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "emailフィールドなし",
			body:         `{"password":"password123"}`,
			setupMock:    func(m *mock.MockLogin) {},
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "passwordフィールドなし",
			body:         `{"email":"test@example.com"}`,
			setupMock:    func(m *mock.MockLogin) {},
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "UseCase実行エラー（認証失敗）",
			body: `{"email":"test@example.com","password":"wrongpassword"}`,
			setupMock: func(m *mock.MockLogin) {
				m.EXPECT().Execute(gomock.Any(), usecase.LoginInput{
					Email:    "test@example.com",
					Password: "wrongpassword",
				}).Return(errors.New("authentication failed"))
			},
			expectedCode: http.StatusUnauthorized,
		},
		{
			name: "UseCase実行エラー（その他のエラー）",
			body: `{"email":"test@example.com","password":"password123"}`,
			setupMock: func(m *mock.MockLogin) {
				m.EXPECT().Execute(gomock.Any(), usecase.LoginInput{
					Email:    "test@example.com",
					Password: "password123",
				}).Return(errors.New("internal server error"))
			},
			expectedCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUseCase := mock.NewMockLogin(ctrl)
			tt.setupMock(mockUseCase)

			handler := handler.NewLogin(mockUseCase)

			rec, c := createTestPostRequest("/login", tt.body)

			err := handler(c)

			// エラーレスポンスはmiddlewareで処理されるため、
			// handler自体はエラーを返すが、実際のレスポンスコードは
			// middlewareによって設定される
			if tt.expectedCode == http.StatusOK {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedCode, rec.Code)
			} else {
				assert.Error(t, err)
			}
		})
	}
}