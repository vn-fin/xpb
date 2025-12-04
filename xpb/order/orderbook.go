package order

import "time"

// OrderBookInfo represents an order book update message.
type OrderBookInfo struct {
	Symbol    string    `json:"symbol"`
	TimeF     float64   `json:"time"`
	TimeT     time.Time `json:"_"`
	BidPrices []float64 `json:"bp"`
	AskPrices []float64 `json:"ap"`
	BidSizes  []float64 `json:"bq"`
	AskSizes  []float64 `json:"aq"`
}

func (c *OrderBookInfo) Build() {
	// Parse time from int64 to time.Time
	seconds := int64(c.TimeF)
	nanos := int64((c.TimeF - float64(seconds)) * 1e9)
	c.TimeT = time.Unix(seconds, nanos)
}
