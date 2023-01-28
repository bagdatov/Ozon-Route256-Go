package repository

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"gitlab.ozon.dev/bagdatov/homework-3/service-store/internal"
	"gitlab.ozon.dev/bagdatov/homework-3/service-store/internal/entity"
)

type database struct {
	pgx *pgxpool.Pool
}

func NewDB(pool *pgxpool.Pool) *database {
	return &database{
		pgx: pool,
	}
}

func (db *database) Add(ctx context.Context, item entity.Item) (uint64, error) {

	var id uint64

	err := db.pgx.QueryRow(ctx,
		`INSERT INTO
			items (seller_id, price, item_name, creation_date)
			VALUES ($1, $2, $3, $4)
			RETURNING id`,

		item.SellerID, item.Price, item.Name, item.Date,
	).Scan(&id)

	return id, err
}

func (db *database) Fetch(ctx context.Context, itemID uint64) (entity.Item, error) {

	var item entity.Item

	err := db.pgx.QueryRow(ctx,
		`SELECT
			id, seller_id, price, item_name, creation_date
			FROM items
			WHERE id = $1`,

		itemID,
	).Scan(&item.ItemID, &item.SellerID, &item.Price, &item.Name, &item.Date)

	if errors.Is(err, pgx.ErrNoRows) {
		return item, internal.NotFound
	}

	return item, err
}

func (db *database) AddReservation(ctx context.Context, order entity.Order) (uint64, error) {

	var id uint64

	err := db.pgx.QueryRow(ctx,
		`INSERT INTO 
			reservations (item_id, order_id, update_date)
			VALUES ($1, $2, $3)
			RETURNING id`,
		order.ItemID, order.ID, order.Date,
	).Scan(&id)

	return id, err
}

func (db *database) CancelReservation(ctx context.Context, order entity.CancelOrder) error {

	_, err := db.pgx.Exec(ctx,
		`DELETE FROM 
			reservations
			WHERE order_id = $1`,

		order.OrderID,
	)
	return err
}

func (db *database) FetchReservation(ctx context.Context, itemID uint64) (entity.Reservation, error) {

	var res entity.Reservation

	err := db.pgx.QueryRow(ctx,
		`SELECT 
			id, item_id, order_id, update_date
			FROM reservations
			WHERE item_id = $1`,

		itemID,
	).Scan(
		&res.ReservationID,
		&res.ItemID,
		&res.OrderID,
		&res.UpdateDate,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return res, internal.NotFound
	}

	return res, err
}
