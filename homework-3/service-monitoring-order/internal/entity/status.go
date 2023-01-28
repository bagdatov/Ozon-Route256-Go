package entity

import "time"

type OrderStatus struct {
	OrderID     uint64    `json:"order_id"`
	ItemID      uint64    `json:"item_id"`
	SellerID    uint64    `json:"seller_id"`
	ClientID    uint64    `json:"client_id"`
	Step        string    `json:"step"`
	IsReserved  bool      `json:"is_reserved"`
	IsCancelled bool      `json:"is_cancelled"`
	UpdateDate  time.Time `json:"update_date"`
}
