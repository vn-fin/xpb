package order

import "time"

// TickInfo represents a trade tick message.
type TickInfo struct {
	Symbol string    `json:"symbol"`
	TimeF  float64   `json:"time"`
	TimeT  time.Time `json:"_"`
	Price  float64   `json:"price"`
	Volume float64   `json:"vol"`
	Side   string    `json:"side"`
}

func (c *TickInfo) Build() {
	seconds := int64(c.TimeF)
	nanos := int64((c.TimeF - float64(seconds)) * 1e9)
	c.TimeT = time.Unix(seconds, nanos)
}
