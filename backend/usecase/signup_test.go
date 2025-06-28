package usecase_test

import (
	"backend/domain/model"
	"backend/infra/persistence"
	mock "backend/test/mock/repository"
	"backend/usecase"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestSignUp_Execute(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name      string
		input     usecase.SignUpInput
		setupMock func(*mock.MockAccount, *mock.MockUser, *mock.MockTransaction)
		wantErr   error
	}{
		{
			name: "正常なサインアップ",
			input: usecase.SignUpInput{
				UID:             "firebase-uid-123",
				Email:           "test@example.com",
				Name:            "test-name",
				ProfileImageURL: "",
			},
			setupMock: func(acc *mock.MockAccount, user *mock.MockUser, tx *mock.MockTransaction) {
				// 期待されるUser作成
				expectedUser, _ := model.NewUser("firebase-uid-123", "test@example.com")

				// トランザクション内の処理
				tx.EXPECT().WithTx(ctx, gomock.AssignableToTypeOf(func(context.Context) error { return nil })).DoAndReturn(
					func(ctx context.Context, fn func(context.Context) error) error {
						// ユーザー作成
						user.EXPECT().Store(ctx, &expectedUser).Return(nil)
						// アカウント作成（AccountIDが動的生成のため、カスタムマッチャーが必要）
						acc.EXPECT().Store(ctx, gomock.AssignableToTypeOf(&model.Account{})).Return(nil)
						return fn(ctx)
					},
				)
			},
			wantErr: nil,
		},
		{
			name: "無効なユーザーID",
			input: usecase.SignUpInput{
				UID:             "",
				Email:           "test@example.com",
				Name:            "test-name",
				ProfileImageURL: "",
			},
			setupMock: func(acc *mock.MockAccount, user *mock.MockUser, tx *mock.MockTransaction) {
				// WithTxが呼び出され、内部でmodel.NewUserがバリデーションエラーを返す
				tx.EXPECT().WithTx(ctx, gomock.AssignableToTypeOf(func(context.Context) error { return nil })).DoAndReturn(
					func(ctx context.Context, fn func(context.Context) error) error {
						return fn(ctx) // NewUserでバリデーションエラーが発生
					},
				)
			},
			wantErr: errors.New("id is required"),
		},
		{
			name: "空のアカウント名（ドメインバリデーションエラー）",
			input: usecase.SignUpInput{
				UID:             "firebase-uid-123",
				Email:           "test@example.com",
				Name:            "",
				ProfileImageURL: "",
			},
			setupMock: func(acc *mock.MockAccount, user *mock.MockUser, tx *mock.MockTransaction) {
				expectedUser, _ := model.NewUser("firebase-uid-123", "test@example.com")

				// トランザクション内でドメインバリデーションエラーが発生
				tx.EXPECT().WithTx(ctx, gomock.AssignableToTypeOf(func(context.Context) error { return nil })).DoAndReturn(
					func(ctx context.Context, fn func(context.Context) error) error {
						// User.Storeは成功するが、Account作成でバリデーションエラーが発生
						user.EXPECT().Store(ctx, &expectedUser).Return(nil)
						return fn(ctx) // NewAccountでバリデーションエラーが発生
					},
				)
			},
			wantErr: errors.New("name is required"), // model.NewAccountで返されるエラー
		},
		{
			name: "ユーザー保存エラー",
			input: usecase.SignUpInput{
				UID:             "firebase-uid-123",
				Email:           "test@example.com",
				Name:            "test-name",
				ProfileImageURL: "",
			},
			setupMock: func(acc *mock.MockAccount, user *mock.MockUser, tx *mock.MockTransaction) {
				expectedUser, _ := model.NewUser("firebase-uid-123", "test@example.com")

				tx.EXPECT().WithTx(ctx, gomock.AssignableToTypeOf(func(context.Context) error { return nil })).DoAndReturn(
					func(ctx context.Context, fn func(context.Context) error) error {
						user.EXPECT().Store(ctx, &expectedUser).Return(persistence.ErrDuplicateRecord)
						return fn(ctx)
					},
				)
			},
			wantErr: persistence.ErrDuplicateRecord,
		},
		{
			name: "アカウント保存エラー",
			input: usecase.SignUpInput{
				UID:             "firebase-uid-123",
				Email:           "test@example.com",
				Name:            "test-name",
				ProfileImageURL: "",
			},
			setupMock: func(acc *mock.MockAccount, user *mock.MockUser, tx *mock.MockTransaction) {
				expectedUser, _ := model.NewUser("firebase-uid-123", "test@example.com")

				tx.EXPECT().WithTx(ctx, gomock.AssignableToTypeOf(func(context.Context) error { return nil })).DoAndReturn(
					func(ctx context.Context, fn func(context.Context) error) error {
						user.EXPECT().Store(ctx, &expectedUser).Return(nil)
						acc.EXPECT().Store(ctx, gomock.AssignableToTypeOf(&model.Account{})).Return(persistence.ErrDuplicateRecord)
						return fn(ctx)
					},
				)
			},
			wantErr: persistence.ErrDuplicateRecord,
		},
		{
			name: "トランザクションエラー",
			input: usecase.SignUpInput{
				UID:             "firebase-uid-123",
				Email:           "test@example.com",
				Name:            "test-name",
				ProfileImageURL: "",
			},
			setupMock: func(acc *mock.MockAccount, user *mock.MockUser, tx *mock.MockTransaction) {
				tx.EXPECT().WithTx(ctx, gomock.AssignableToTypeOf(func(context.Context) error { return nil })).Return(errors.New("transaction failed"))
			},
			wantErr: errors.New("transaction failed"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockAccount := mock.NewMockAccount(ctrl)
			mockUser := mock.NewMockUser(ctrl)
			mockTx := mock.NewMockTransaction(ctrl)

			signUp := usecase.NewSignUp(mockAccount, mockUser, mockTx)

			tt.setupMock(mockAccount, mockUser, mockTx)

			err := signUp.Execute(context.Background(), tt.input)

			if tt.wantErr != nil {
				assert.Error(t, err)
				// エラーメッセージの部分一致で確認（ErrorIsはAppErrorのラップによりうまく動作しない場合がある）
				assert.Contains(t, err.Error(), tt.wantErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
