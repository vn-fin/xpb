package x

import "time"

// TickInfo represents a trade tick message.
type TickInfo struct {
	Symbol  string    `json:"symbol"`
	TimeInt int64     `json:"time"`
	TimeT   time.Time `json:"_"`
	Price   float64   `json:"price"`
	Volume  float64   `json:"vol"`
	Size    float64   `json:"size"`
}

func (c *TickInfo) Build() {
	// Parse time from int64 to time.Time
	c.TimeT = time.Unix(c.TimeInt, 0)
}
