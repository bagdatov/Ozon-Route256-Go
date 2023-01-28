package usecase

import (
	"context"
	"gitlab.ozon.dev/bagdatov/homework-3/service-store/internal/entity"
)

type MessageBroker interface {
	Publish(reservation entity.Reservation) error
	Cancel(order entity.CancelOrder) error
}

type StorageRepo interface {
	Add(ctx context.Context, item entity.Item) (id uint64, err error)
	Fetch(ctx context.Context, itemID uint64) (entity.Item, error)
}

type CacheRepo interface {
	Get(ctx context.Context, ID uint64) (item entity.Item, err error)
	Set(ctx context.Context, item entity.Item) error
	Delete(ctx context.Context, itemID uint64) error
}

type ReservationRepo interface {
	AddReservation(ctx context.Context, order entity.Order) (id uint64, err error)
	CancelReservation(ctx context.Context, order entity.CancelOrder) error
	FetchReservation(ctx context.Context, itemID uint64) (entity.Reservation, error)
}

type Storage interface {
	AddItem(ctx context.Context, item entity.Item) (id uint64, err error)
	FetchItem(ctx context.Context, itemID uint64) (entity.Item, error)
	ReserveItem(ctx context.Context, order entity.Order) error
	CancelReservation(ctx context.Context, order entity.CancelOrder) error
}
