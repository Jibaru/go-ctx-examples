package mysql

import (
	"context"
	"database/sql"

	"github.com/jibaru/ctx-transaction/internal/orders/domain"
	"github.com/jibaru/ctx-transaction/internal/shared/app"
)

type MySQLOrderRepository struct {
	db *sql.DB
}

func NewMySQLOrderRepository(db *sql.DB) *MySQLOrderRepository {
	return &MySQLOrderRepository{db: db}
}

func (r *MySQLOrderRepository) Save(ctx context.Context, order *domain.Order) error {
	tx, ok := ctx.Value(app.SessionKey).(*sql.Tx)
	if ok {
		_, err := tx.Exec("INSERT INTO orders (id, customer_id) VALUES (?, ?)", order.ID, order.CustomerID)
		return err
	}

	_, err := r.db.Exec("INSERT INTO orders (id, customer_id) VALUES (?, ?)", order.ID, order.CustomerID)
	return err
}
