package mysql

import (
	"context"
	"database/sql"

	"github.com/jibaru/ctx-transaction/internal/orders/domain"
	"github.com/jibaru/ctx-transaction/internal/shared/app"
)

type MySQLOrderRepository struct {
	commonRepository
	db *sql.DB
}

func NewMySQLOrderRepository(db *sql.DB) *MySQLOrderRepository {
	return &MySQLOrderRepository{db: db}
}

func (r *MySQLOrderRepository) Save(ctx context.Context, order *domain.Order) error {
	query := "INSERT INTO orders (id, customer_name, description created_on) VALUES (?, ?, ?, ?)"

	tx, ok := ctx.Value(app.SessionKey).(*sql.Tx)
	if ok {
		_, err := tx.Exec(query, order.ID, order.CustomerName, order.Description, order.CreatedOn)
		return err
	}

	_, err := r.db.Exec(query, order.ID, order.CustomerName, order.Description, order.CreatedOn)
	return err
}
