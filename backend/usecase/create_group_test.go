package usecase_test

import (
	"backend/domain/model"
	mock "backend/test/mock/repository"
	"backend/usecase"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateGroup(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	accountRepo := mock.NewMockAccount(ctrl)
	groupRepo := mock.NewMockGroup(ctrl)

	tests := []struct {
		name    string
		input   usecase.CreateGroupInput
		setup   func()
		wantErr bool
		errMsg  string
	}{
		{
			name: "正常系：新しいグループを作成",
			input: usecase.CreateGroupInput{
				UserID:      "test-user-id",
				Name:        "テストグループ",
				Description: "テスト用のグループです",
			},
			setup: func() {
				account := model.Account{
					ID:     model.NewAccountID(),
					UserID: "test-user-id",
					Name:   "テストユーザー",
				}
				accountRepo.EXPECT().FindByUserID(gomock.Any(), "test-user-id").Return(account, nil)
				groupRepo.EXPECT().Store(gomock.Any(), gomock.Any()).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "正常系：説明なしでグループを作成",
			input: usecase.CreateGroupInput{
				UserID:      "test-user-id",
				Name:        "説明なしグループ",
				Description: "",
			},
			setup: func() {
				account := model.Account{
					ID:     model.NewAccountID(),
					UserID: "test-user-id",
					Name:   "テストユーザー",
				}
				accountRepo.EXPECT().FindByUserID(gomock.Any(), "test-user-id").Return(account, nil)
				groupRepo.EXPECT().Store(gomock.Any(), gomock.Any()).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "異常系：存在しないユーザーID",
			input: usecase.CreateGroupInput{
				UserID:      "non-existent-user",
				Name:        "テストグループ",
				Description: "説明",
			},
			setup: func() {
				accountRepo.EXPECT().FindByUserID(gomock.Any(), "non-existent-user").Return(model.Account{}, assert.AnError)
			},
			wantErr: true,
		},
		{
			name: "異常系：グループ名が空文字",
			input: usecase.CreateGroupInput{
				UserID:      "test-user-id",
				Name:        "",
				Description: "説明があってもグループ名が空",
			},
			setup: func() {
				account := model.Account{
					ID:     model.NewAccountID(),
					UserID: "test-user-id",
					Name:   "テストユーザー",
				}
				accountRepo.EXPECT().FindByUserID(gomock.Any(), "test-user-id").Return(account, nil)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()

			createGroup := usecase.NewCreateGroup(accountRepo, groupRepo)
			err := createGroup.Execute(context.Background(), tt.input)

			if tt.wantErr {
				assert.Error(t, err)
				if tt.errMsg != "" {
					assert.Contains(t, err.Error(), tt.errMsg)
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}