package usecase

import (
	"context"
	"gitlab.ozon.dev/bagdatov/homework-3/service-create-order/internal/entity"
)

type OrderRepo interface {
	Create(ctx context.Context, order entity.Order) (id uint64, err error)
	Fetch(ctx context.Context, id uint64) (entity.Order, error)
	Cancel(ctx context.Context, order entity.CancelOrder) error
}

type Order interface {
	Create(ctx context.Context, order entity.Order) (id uint64, err error)
	Cancel(ctx context.Context, id uint64) error
}

type MessageBroker interface {
	Publish(order entity.Order) error
	Cancel(order entity.CancelOrder) error
}
