package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"gitlab.ozon.dev/bagdatov/homework-3/service-monitoring-order/internal"
	"gitlab.ozon.dev/bagdatov/homework-3/service-monitoring-order/internal/entity"
)

type monitoringUseCase struct {
	res ReservationRepo
	st  StatusRepo
	br  MessageBroker
}

func New(res ReservationRepo, st StatusRepo, br MessageBroker) *monitoringUseCase {
	return &monitoringUseCase{
		res: res,
		st:  st,
		br:  br,
	}
}

const (
	RESERVATION = "UNDER RESERVATION"
	SHIPPING    = "SHIPPING"
	CANCELLED   = "CANCELLED"
)

// prometheus metrics
var (
	incomingOrdersCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "incoming_order",
		Help: "The total number of started orders",
	})
	reservationConfirmationCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "reservation_order",
		Help: "The total number of order reservation confirmations",
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

func (m *monitoringUseCase) FetchStatus(ctx context.Context, orderID uint64) (entity.OrderStatus, error) {

	status, err := m.st.FetchStatus(ctx, orderID)
	if err != nil {
		usecaseErrorsCounter.Inc()
		return status, fmt.Errorf("monitoringUseCase.FetchStatus - st.FetchStatus: %w", err)
	}

	if status.IsCancelled {
		status.Step = CANCELLED
		return status, nil
	}

	reservation, err := m.res.FetchReservation(ctx, orderID)
	if errors.Is(err, internal.NotFound) {
		status.Step = RESERVATION
		return status, nil

	} else if err != nil {
		usecaseErrorsCounter.Inc()
		return status, fmt.Errorf("monitoringUseCase.FetchStatus - res.FetchReservation: %w", err)
	}

	status.UpdateDate = reservation.UpdateDate
	status.IsReserved = true
	status.Step = SHIPPING

	return status, nil
}

func (m *monitoringUseCase) AddOrder(ctx context.Context, order entity.Order) error {

	status := entity.OrderStatus{
		OrderID:    order.ID,
		ItemID:     order.ItemID,
		SellerID:   order.SellerID,
		ClientID:   order.ClientID,
		UpdateDate: order.Date,
	}

	incomingOrdersCounter.Inc()
	err := m.st.AddStatus(ctx, status)
	if err != nil {
		usecaseErrorsCounter.Inc()
		return fmt.Errorf("monitoringUseCase.AddOrder - res.AddStatus: %w", err)
	}

	return nil
}

func (m *monitoringUseCase) MarkReservation(ctx context.Context, reserve entity.Reservation) error {
	reservationConfirmationCounter.Inc()
	err := m.res.AddReservation(ctx, reserve)
	if err != nil {
		usecaseErrorsCounter.Inc()
		return fmt.Errorf("monitoringUseCase.MarkReservation - res.AddReservation: %w", err)
	}
	return nil
}

func (m *monitoringUseCase) Cancel(ctx context.Context, orderID uint64) error {

	cancelOrdersCounter.Inc()
	err := m.st.Cancel(ctx, orderID)
	if err != nil {
		usecaseErrorsCounter.Inc()
		return fmt.Errorf("monitoringUseCase.Cancel - st.Cancel: %w", err)
	}

	err = m.res.DeleteReservation(ctx, orderID)
	if err != nil {
		usecaseErrorsCounter.Inc()
		return fmt.Errorf("monitoringUseCase.Cancel - res.DeleteReservation: %w", err)
	}

	return nil
}
