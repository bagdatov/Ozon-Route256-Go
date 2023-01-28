package entity

import "time"

type Item struct {
	ItemID     uint64    `json:"item_id"`
	SellerID   uint64    `json:"seller_id"`
	Price      uint64    `json:"price"`
	Name       string    `json:"item_name"`
	IsReserved bool      `json:"is_reserved"`
	Date       time.Time `json:"date"`
}
