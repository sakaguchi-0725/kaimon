package persistence

import (
	"backend/core"
	"backend/domain/repository"
	"backend/infra/db"
	"context"
	"fmt"
)

type transaction struct {
	conn *db.Conn
}

func (t *transaction) WithTx(ctx context.Context, fn func(ctx context.Context) error) error {
	tx := t.conn.Begin()
	if err := tx.Error; err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	txConn := &db.Conn{DB: tx}
	ctx = context.WithValue(ctx, core.TxKey, txConn)

	if err := fn(ctx); err != nil {
		if rbErr := tx.Rollback().Error; rbErr != nil {
			return fmt.Errorf("rollback failed after error: %v, rollback error: %w", err, rbErr)
		}
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("failed to commit: %w", err)
	}

	return nil
}

func GetTx(ctx context.Context) *db.Conn {
	if tx, ok := ctx.Value(core.TxKey).(*db.Conn); ok {
		return tx
	}
	return nil
}

func NewTransaction(c *db.Conn) repository.Transaction {
	return &transaction{conn: c}
}
