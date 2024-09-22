package mysql

import (
	"context"
	"database/sql"

	"github.com/jibaru/ctx-transaction/internal/orders/domain"
	"github.com/jibaru/ctx-transaction/internal/shared/app"
)

type MySQLOrderLineRepository struct {
	db *sql.DB
}

func NewMySQLOrderLineRepository(db *sql.DB) *MySQLOrderLineRepository {
	return &MySQLOrderLineRepository{db: db}
}

func (r *MySQLOrderLineRepository) Save(ctx context.Context, orderLine *domain.OrderLine) error {
	tx, ok := ctx.Value(app.SessionKey).(*sql.Tx)
	if ok {
		_, err := tx.Exec("INSERT INTO order_lines (id, order_id, name, quantity) VALUES (?, ?, ?, ?)", orderLine.ID, orderLine.OrderID, orderLine.Name, orderLine.Quantity)
		return err
	}

	_, err := r.db.Exec("INSERT INTO order_lines (id, order_id, name, quantity) VALUES (?, ?, ?, ?)", orderLine.ID, orderLine.OrderID, orderLine.Name, orderLine.Quantity)
	return err
}
