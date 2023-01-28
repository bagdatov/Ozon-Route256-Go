package repository

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
	"gitlab.ozon.dev/bagdatov/homework-3/service-monitoring-order/internal"

	"github.com/jackc/pgx/v4/pgxpool"
	"gitlab.ozon.dev/bagdatov/homework-3/service-monitoring-order/internal/entity"
)

type database struct {
	pgx *pgxpool.Pool
}

func NewDB(pool *pgxpool.Pool) *database {
	return &database{
		pgx: pool,
	}
}

func (db *database) AddStatus(ctx context.Context, order entity.OrderStatus) error {

	_, err := db.pgx.Exec(ctx,
		`INSERT INTO 
			order_status (order_id, item_id, seller_id, client_id, update_date, is_cancelled)
			VALUES ($1, $2, $3, $4, $5, false)`,

		order.OrderID,
		order.ItemID,
		order.SellerID,
		order.ClientID,
		order.UpdateDate,
	)

	return err
}

func (db *database) FetchStatus(ctx context.Context, orderID uint64) (entity.OrderStatus, error) {

	var order entity.OrderStatus

	err := db.pgx.QueryRow(ctx,
		`SELECT
		order_id, item_id, seller_id, client_id, update_date, is_cancelled
		FROM order_status
		WHERE order_id = $1`,

		orderID,
	).Scan(
		&order.OrderID,
		&order.ItemID,
		&order.SellerID,
		&order.ClientID,
		&order.UpdateDate,
		&order.IsCancelled)

	if errors.Is(err, pgx.ErrNoRows) {
		return order, internal.NotFound
	}

	return order, err
}

func (db *database) Cancel(ctx context.Context, orderID uint64) error {

	_, err := db.pgx.Exec(ctx,
		`UPDATE order_status
			SET is_cancelled = true
			WHERE order_id = $1`,
		orderID,
	)

	return err
}

func (db *database) AddReservation(ctx context.Context, reserve entity.Reservation) error {

	_, err := db.pgx.Exec(ctx,
		`INSERT INTO 
			reservations (id, order_id, update_date)
			VALUES ($1, $2, $3)`,

		reserve.ReservationID,
		reserve.OrderID,
		reserve.UpdateDate,
	)

	return err
}

func (db *database) FetchReservation(ctx context.Context, orderID uint64) (entity.Reservation, error) {

	var reservation entity.Reservation

	err := db.pgx.QueryRow(ctx,
		`SELECT
		id, order_id, update_date
		FROM reservations
		WHERE order_id = $1`,

		orderID,
	).Scan(
		&reservation.ReservationID,
		&reservation.OrderID,
		&reservation.UpdateDate,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return reservation, internal.NotFound
	}

	return reservation, err
}

func (db *database) DeleteReservation(ctx context.Context, orderID uint64) error {

	_, err := db.pgx.Exec(ctx,
		`DELETE FROM 
			reservations
			WHERE order_id = $1`,
		orderID,
	)

	return err
}
