package usecase

import (
	"context"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"gitlab.ozon.dev/bagdatov/homework-3/service-create-order/internal/entity"
	"time"
)

type orderUseCase struct {
	repo   OrderRepo
	broker MessageBroker
}

// prometheus metrics
var (
	startOrdersCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "create_order",
		Help: "The total number of start orders",
	})
	cancelOrdersCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "cancel_order",
		Help: "The total number of order cancellations",
	})
	usecaseErrorsCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "usecase_error",
		Help: "The total number of errors",
	})
)

const CancelReason = "incoming cancellation request"

func New(repo OrderRepo, broker MessageBroker) *orderUseCase {
	return &orderUseCase{
		repo:   repo,
		broker: broker,
	}
}

func (uc *orderUseCase) Create(ctx context.Context, order entity.Order) (uint64, error) {
	startOrdersCounter.Inc()
	id, err := uc.repo.Create(ctx, order)
	if err != nil {
		usecaseErrorsCounter.Inc()
		return 0, fmt.Errorf("orderUseCase.Create - repo.Create: %w", err)
	}

	order.ID = id

	if err := uc.broker.Publish(order); err != nil {
		usecaseErrorsCounter.Inc()
		return 0, fmt.Errorf("orderUseCase.Create - broker.Publish: %w", err)
	}

	return id, nil
}

func (uc *orderUseCase) Cancel(ctx context.Context, id uint64) error {
	cancelOrdersCounter.Inc()
	order, err := uc.repo.Fetch(ctx, id)
	if err != nil {
		usecaseErrorsCounter.Inc()
		return fmt.Errorf("orderUseCase.Cancel - repo.Fetch: %w", err)
	}

	request := entity.CancelOrder{
		OrderID:  order.ID,
		ItemID:   order.ItemID,
		SellerID: order.SellerID,
		ClientID: order.ClientID,
		Reason:   CancelReason,
		Date:     time.Now(),
	}

	if err := uc.repo.Cancel(ctx, request); err != nil {
		usecaseErrorsCounter.Inc()
		return fmt.Errorf("orderUseCase.Cancel - repo.Cancel: %w", err)
	}

	if err := uc.broker.Cancel(request); err != nil {
		usecaseErrorsCounter.Inc()
		return fmt.Errorf("orderUseCase.Cancel - broker.Cancel: %w", err)
	}

	return nil
}
