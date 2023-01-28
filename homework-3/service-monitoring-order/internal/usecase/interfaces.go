package usecase

import (
	"context"
	"gitlab.ozon.dev/bagdatov/homework-3/service-monitoring-order/internal/entity"
)

type StatusRepo interface {
	AddStatus(ctx context.Context, order entity.OrderStatus) error
	FetchStatus(ctx context.Context, orderID uint64) (entity.OrderStatus, error)
	Cancel(ctx context.Context, orderID uint64) error
}

type ReservationRepo interface {
	AddReservation(ctx context.Context, reserve entity.Reservation) error
	FetchReservation(ctx context.Context, orderID uint64) (entity.Reservation, error)
	DeleteReservation(ctx context.Context, orderID uint64) error
}

type Monitoring interface {
	AddOrder(ctx context.Context, order entity.Order) error
	MarkReservation(ctx context.Context, reserve entity.Reservation) error
	FetchStatus(ctx context.Context, orderID uint64) (entity.OrderStatus, error)
	Cancel(ctx context.Context, orderID uint64) error
}

type MessageBroker interface {
	Cancel(order entity.CancelOrder) error
}
