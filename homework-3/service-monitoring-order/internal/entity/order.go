package entity

import (
	"time"
)

type Order struct {
	ID       uint64    `json:"id"`
	ItemID   uint64    `json:"item_id"`
	SellerID uint64    `json:"seller_id"`
	ClientID uint64    `json:"client_id"`
	Date     time.Time `json:"date"`
}
