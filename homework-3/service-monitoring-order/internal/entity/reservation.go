package entity

import "time"

type Reservation struct {
	ReservationID uint64    `json:"reservation_id"`
	OrderID       uint64    `json:"order_id"`
	UpdateDate    time.Time `json:"update_date"`
}
