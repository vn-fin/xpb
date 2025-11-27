package x

import "time"

type OrderInfo struct {
	OrderID      string    `json:"order_id"`
	Symbol       string    `json:"symbol"`
	Side         SideT     `json:"side"`
	OrderType    TypeT     `json:"order_type"`
	OrderPrice   float64   `json:"order_price"`
	Quantity     float64   `json:"quantity"`
	FilledQty    float64   `json:"filled_qty"`
	RemainingQty float64   `json:"remaining_qty"`
	Status       StatusT   `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
