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

func TestSignUp(t *testing.T) {
	tests := []struct {
		name         string
		authHeader   string
		setupMock    func(*mock.MockSignUp)
		expectedCode int
	}{
		{
			name:       "正常なリクエスト",
			authHeader: "Bearer valid-token",
			setupMock: func(m *mock.MockSignUp) {
				m.EXPECT().Execute(gomock.Any(), usecase.SignUpInput{
					IDToken: "valid-token",
				}).Return(nil)
			},
			expectedCode: http.StatusNoContent,
		},
		{
			name:         "Authorizationヘッダーなし",
			authHeader:   "",
			setupMock:    func(m *mock.MockSignUp) {},
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "Bearer プレフィックスなし",
			authHeader:   "invalid-token",
			setupMock:    func(m *mock.MockSignUp) {},
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "空のトークン",
			authHeader:   "Bearer ",
			setupMock:    func(m *mock.MockSignUp) {},
			expectedCode: http.StatusBadRequest,
		},
		{
			name:       "UseCase実行エラー",
			authHeader: "Bearer valid-token",
			setupMock: func(m *mock.MockSignUp) {
				m.EXPECT().Execute(gomock.Any(), usecase.SignUpInput{
					IDToken: "valid-token",
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

			rec, c := createTestPostRequest("/signup", "")
			if tt.authHeader != "" {
				c.Request().Header.Set("Authorization", tt.authHeader)
			}

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
