package x

// SideT presents the side of an order: buy or sell
type SideT string

// TypeT presents the type of an order: market, limit, ATO, ATC
// We do not implement other types such as stop loss, stop limit, etc.
type TypeT string

// StatusT presents the status of an order
type StatusT int

const (
	SideBuy  SideT = "B"
	SideSell SideT = "S"
)

const (
	TypeMarket TypeT = "M"
	TypeLimit  TypeT = "L"
	TypeATO    TypeT = "ATO"
	TypeATC    TypeT = "ATC"
)

const (
	StatusCreated         StatusT = 1
	StatusPending         StatusT = 2
	StatusFilled          StatusT = 3
	StatusCanceled        StatusT = 4
	StatusRejected        StatusT = 5
	StatusExpired         StatusT = 6
	StatusPartiallyFilled StatusT = 7
)
