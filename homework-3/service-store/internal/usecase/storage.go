package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"gitlab.ozon.dev/bagdatov/homework-3/service-store/internal"
	"gitlab.ozon.dev/bagdatov/homework-3/service-store/internal/entity"
	"log"
	"time"
)

type storageUseCase struct {
	res    ReservationRepo
	store  StorageRepo
	broker MessageBroker
	cache  CacheRepo
}

func New(res ReservationRepo, store StorageRepo, broker MessageBroker, cache CacheRepo) *storageUseCase {
	return &storageUseCase{
		res:    res,
		store:  store,
		broker: broker,
		cache:  cache,
	}
}

const (
	ReservationFailure = "cannot reserve item"
)

// prometheus metrics
var (
	addItemCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "add_item",
		Help: "The total number of start orders",
	})
	reservationCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "reservation_order",
		Help: "The total number of reservation requests",
	})

	failedReservationCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "reservation_order_failure",
		Help: "The total number of failed reservation requests",
	})
	cancelOrdersCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "cancel_reservation",
		Help: "The total number of order cancellations",
	})
	usecaseErrorsCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "usecase_error",
		Help: "The total number of errors",
	})
)

func (uc *storageUseCase) AddItem(ctx context.Context, item entity.Item) (uint64, error) {
	addItemCounter.Inc()
	id, err := uc.store.Add(ctx, item)
	if err != nil {
		usecaseErrorsCounter.Inc()
		return id, fmt.Errorf("storageUseCase.AddItem - store.Add error: %w", err)
	}

	return id, nil
}

func (uc *storageUseCase) FetchItem(ctx context.Context, itemID uint64) (entity.Item, error) {

	// check in the cache
	item, err := uc.cache.Get(ctx, itemID)

	if errors.Is(err, internal.NotFound) {

		item, err = uc.store.Fetch(ctx, itemID)
		if err != nil {
			usecaseErrorsCounter.Inc()
			return item, fmt.Errorf("storageUseCase.FetchItem - store.Fetch error: %w", err)
		}

		item.IsReserved = true

		_, err = uc.res.FetchReservation(ctx, itemID)
		if errors.Is(err, internal.NotFound) {
			item.IsReserved = false

		} else if err != nil {
			usecaseErrorsCounter.Inc()
			return item, fmt.Errorf("storageUseCase.FetchItem - res.FetchReservation error: %w", err)
		}

		err = uc.cache.Set(ctx, item)
		if err != nil {
			usecaseErrorsCounter.Inc()
			return item, fmt.Errorf("storageUseCase.FetchItem - cache.Set error: %w", err)
		}
	}

	return item, nil
}

func (uc *storageUseCase) ReserveItem(ctx context.Context, order entity.Order) error {
	reservationCounter.Inc()

	id, err := uc.res.AddReservation(ctx, order)
	if err != nil {
		failedReservationCounter.Inc()

		errCancel := uc.broker.Cancel(entity.CancelOrder{
			OrderID:  order.ID,
			ItemID:   order.ItemID,
			SellerID: order.SellerID,
			ClientID: order.ClientID,
			Reason:   ReservationFailure,
			Date:     time.Now(),
		})

		if errCancel != nil {
			usecaseErrorsCounter.Inc()
			log.Printf("cannot send order cancellation error: <%v>, request: %+v", err, order)
		}

		usecaseErrorsCounter.Inc()
		return fmt.Errorf("storageUseCase.ReserveItem - res.AddReservation error: %w", err)
	}
	err = uc.cache.Delete(ctx, order.ItemID)
	if err != nil {
		usecaseErrorsCounter.Inc()
		return fmt.Errorf("storageUseCase.ReserveItem - cache.Delete error: %w", err)
	}

	return uc.broker.Publish(entity.Reservation{
		ReservationID: id,
		ItemID:        order.ItemID,
		OrderID:       order.ID,
		UpdateDate:    time.Now(),
	})
}

func (uc *storageUseCase) CancelReservation(ctx context.Context, order entity.CancelOrder) error {

	cancelOrdersCounter.Inc()

	err := uc.res.CancelReservation(ctx, order)
	if err != nil {
		usecaseErrorsCounter.Inc()
		return fmt.Errorf("storageUseCase.CancelReservation - res.CancelReservation error: %w", err)
	}

	err = uc.cache.Delete(ctx, order.ItemID)
	if err != nil {
		usecaseErrorsCounter.Inc()
		return fmt.Errorf("storageUseCase.CancelReservation - cache.Delete error: %w", err)
	}

	return nil
}
