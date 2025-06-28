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
	tests := []struct {
		name      string
		input     usecase.SignUpInput
		setupMock func(*mock.MockAuthenticator, *mock.MockAccount, *mock.MockUser, *mock.MockTransaction)
		wantErr   error
	}{
		{
			name: "正常なサインアップ",
			input: usecase.SignUpInput{
				IDToken: "valid_token",
			},
			setupMock: func(auth *mock.MockAuthenticator, acc *mock.MockAccount, user *mock.MockUser, tx *mock.MockTransaction) {
				ctx := context.Background()

				// 認証トークン検証
				auth.EXPECT().VerifyToken("valid_token").Return("firebase-uid-123", "test@example.com", "test-name", nil)

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
			name: "無効なトークン",
			input: usecase.SignUpInput{
				IDToken: "invalid_token",
			},
			setupMock: func(auth *mock.MockAuthenticator, acc *mock.MockAccount, user *mock.MockUser, tx *mock.MockTransaction) {
				auth.EXPECT().VerifyToken("invalid_token").Return("", "", "", persistence.ErrInvalidToken)
				// エラーの場合、Store系やトランザクションの呼び出しは期待しない
			},
			wantErr: persistence.ErrInvalidToken,
		},
		{
			name: "空のアカウント名（ドメインバリデーションエラー）",
			input: usecase.SignUpInput{
				IDToken: "valid_token",
			},
			setupMock: func(auth *mock.MockAuthenticator, acc *mock.MockAccount, user *mock.MockUser, tx *mock.MockTransaction) {
				ctx := context.Background()
				expectedUser, _ := model.NewUser("firebase-uid-123", "test@example.com")

				// 認証トークン検証
				auth.EXPECT().VerifyToken("valid_token").Return("firebase-uid-123", "test@example.com", "", nil)

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
				IDToken: "valid_token",
			},
			setupMock: func(auth *mock.MockAuthenticator, acc *mock.MockAccount, user *mock.MockUser, tx *mock.MockTransaction) {
				ctx := context.Background()
				expectedUser, _ := model.NewUser("firebase-uid-123", "test@example.com")

				auth.EXPECT().VerifyToken("valid_token").Return("firebase-uid-123", "test@example.com", "test-name", nil)

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
				IDToken: "valid_token",
			},
			setupMock: func(auth *mock.MockAuthenticator, acc *mock.MockAccount, user *mock.MockUser, tx *mock.MockTransaction) {
				ctx := context.Background()
				expectedUser, _ := model.NewUser("firebase-uid-123", "test@example.com")

				auth.EXPECT().VerifyToken("valid_token").Return("firebase-uid-123", "test@example.com", "test-name", nil)

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
				IDToken: "valid_token",
			},
			setupMock: func(auth *mock.MockAuthenticator, acc *mock.MockAccount, user *mock.MockUser, tx *mock.MockTransaction) {
				ctx := context.Background()

				auth.EXPECT().VerifyToken("valid_token").Return("firebase-uid-123", "test@example.com", "test-name", nil)
				tx.EXPECT().WithTx(ctx, gomock.AssignableToTypeOf(func(context.Context) error { return nil })).Return(errors.New("transaction failed"))
			},
			wantErr: errors.New("transaction failed"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockAuth := mock.NewMockAuthenticator(ctrl)
			mockAccount := mock.NewMockAccount(ctrl)
			mockUser := mock.NewMockUser(ctrl)
			mockTx := mock.NewMockTransaction(ctrl)

			signUp := usecase.NewSignUp(mockAuth, mockAccount, mockUser, mockTx)

			tt.setupMock(mockAuth, mockAccount, mockUser, mockTx)

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
