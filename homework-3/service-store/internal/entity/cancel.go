package entity

import "time"

type CancelOrder struct {
	OrderID  uint64    `json:"order_id"`
	ItemID   uint64    `json:"item_id"`
	SellerID uint64    `json:"seller_id"`
	ClientID uint64    `json:"client_id"`
	Reason   string    `json:"reason"`
	Date     time.Time `json:"date"`
}
