package repository

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"gitlab.ozon.dev/bagdatov/homework-3/service-create-order/internal"
	"gitlab.ozon.dev/bagdatov/homework-3/service-create-order/internal/entity"
)

type database struct {
	pgx *pgxpool.Pool
}

func NewDB(pool *pgxpool.Pool) *database {
	return &database{
		pgx: pool,
	}
}

func (db *database) Create(ctx context.Context, order entity.Order) (uint64, error) {

	err := db.pgx.QueryRow(ctx,
		`INSERT INTO 
			  orders (item_id, seller_id, client_id, creation_date)
			  VALUES ($1, $2, $3, $4)
			  RETURNING id`,

		order.ItemID,
		order.SellerID,
		order.ClientID,
		order.Date,
	).Scan(&order.ID)

	return order.ID, err
}

func (db *database) Fetch(ctx context.Context, id uint64) (entity.Order, error) {

	var order entity.Order

	err := db.pgx.QueryRow(ctx,
		`SELECT
		id, item_id, seller_id, client_id, creation_date
		FROM orders
		WHERE id = $1`,

		id,
	).Scan(
		&order.ID,
		&order.ItemID,
		&order.SellerID,
		&order.ClientID,
		&order.Date)

	if errors.Is(err, pgx.ErrNoRows) {
		return order, internal.NotFound
	}

	return order, err
}

func (db *database) Cancel(ctx context.Context, order entity.CancelOrder) error {

	_, err := db.pgx.Exec(ctx,
		`INSERT INTO 
			cancel_orders (order_id, cancellation_date)
			VALUES ($1, $2)`,
		order.OrderID,
		order.Date,
	)

	return err
}
