//go:generate mockgen -source=transaction.go -destination=../../test/mock/repository/transaction_mock.go -package=mock
package repository

import (
	"context"
)

type Transaction interface {
	WithTx(ctx context.Context, fn func(ctx context.Context) error) error
}
