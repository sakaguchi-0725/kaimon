package usecase_test

import (
	"backend/domain/model"
	mock "backend/test/mock/repository"
	"backend/usecase"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetAccount_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testUserID := "test-user-id"
	testAccountID := model.NewAccountID()
	testAccount := model.Account{
		ID:     testAccountID,
		UserID: testUserID,
		Name:   "テストユーザー",
	}

	tests := []struct {
		name    string
		input   usecase.GetAccountInput
		setup   func(mockAccountRepo *mock.MockAccount)
		want    usecase.GetAccountOutput
		wantErr error
	}{
		{
			name: "正常にアカウント情報を取得",
			input: usecase.GetAccountInput{
				UserID: testUserID,
			},
			setup: func(mockAccountRepo *mock.MockAccount) {
				mockAccountRepo.EXPECT().
					FindByUserID(gomock.Any(), testUserID).
					Return(testAccount, nil)
			},
			want: usecase.GetAccountOutput{
				ID:   testAccountID.String(),
				Name: "テストユーザー",
			},
			wantErr: nil,
		},
		{
			name: "存在しないユーザーID",
			input: usecase.GetAccountInput{
				UserID: "non-existent-user-id",
			},
			setup: func(mockAccountRepo *mock.MockAccount) {
				mockAccountRepo.EXPECT().
					FindByUserID(gomock.Any(), "non-existent-user-id").
					Return(model.Account{}, errors.New("account not found"))
			},
			want:    usecase.GetAccountOutput{},
			wantErr: errors.New("account not found"),
		},
		{
			name: "空のユーザーID",
			input: usecase.GetAccountInput{
				UserID: "",
			},
			setup: func(mockAccountRepo *mock.MockAccount) {
				mockAccountRepo.EXPECT().
					FindByUserID(gomock.Any(), "").
					Return(model.Account{}, errors.New("invalid user id"))
			},
			want:    usecase.GetAccountOutput{},
			wantErr: errors.New("invalid user id"),
		},
		{
			name: "データベースエラー",
			input: usecase.GetAccountInput{
				UserID: testUserID,
			},
			setup: func(mockAccountRepo *mock.MockAccount) {
				mockAccountRepo.EXPECT().
					FindByUserID(gomock.Any(), testUserID).
					Return(model.Account{}, errors.New("database error"))
			},
			want:    usecase.GetAccountOutput{},
			wantErr: errors.New("database error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockAccountRepo := mock.NewMockAccount(ctrl)
			tt.setup(mockAccountRepo)

			useCase := usecase.NewGetAccount(mockAccountRepo)
			got, err := useCase.Execute(context.Background(), tt.input)

			if tt.wantErr != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.wantErr.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}