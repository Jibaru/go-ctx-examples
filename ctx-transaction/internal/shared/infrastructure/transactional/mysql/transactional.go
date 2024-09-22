package mysql

import (
	"context"
	"database/sql"

	"github.com/jibaru/ctx-transaction/internal/shared/app"
)

type MySQLTransactional struct {
	db *sql.DB
}

func NewMySQLTransactional(db *sql.DB) *MySQLTransactional {
	return &MySQLTransactional{db: db}
}

func (t *MySQLTransactional) InTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	tx, err := t.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	txCtx := context.WithValue(ctx, app.SessionKey, tx)
	if err := fn(txCtx); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
