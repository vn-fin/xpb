package order

import "time"

// OrderBookInfo represents an order book update message.
type OrderBookInfo struct {
	Symbol    string    `json:"symbol"`
	TimeF     float64   `json:"time"`
	TimeT     time.Time `json:"_"`
	BidPrices []float64 `json:"bid_prices"`
	AskPrices []float64 `json:"ask_prices"`
	BidSizes  []float64 `json:"bid_sizes"`
	AskSizes  []float64 `json:"ask_sizes"`
}

func (c *OrderBookInfo) Build() {
	// Parse time from int64 to time.Time
	seconds := int64(c.TimeF)
	nanos := int64((c.TimeF - float64(seconds)) * 1e9)
	c.TimeT = time.Unix(seconds, nanos)
}
