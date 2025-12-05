package order

import "time"

type OrderInfo struct {
	OrderID string `json:"order_id"`
	UserID  string `json:"user_id"`
	//AccountID    string    `json:"account_id"`
	//SubAccountId string    `json:"sub_account_id"`
	CredentialID string    `json:"credential_id"`
	Symbol       string    `json:"symbol"`
	SymbolType   string    `json:"symbol_type"`
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
