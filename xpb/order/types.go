package order

// SideT presents the side of an order: buy or sell
type SideT string

// TypeT presents the type of an order: market, limit, ATO, ATC
// We do not implement other types such as stop loss, stop limit, etc.
type TypeT string

// StatusT presents the status of an order
type StatusT string

const (
	SideBuy  SideT = "B"
	SideSell SideT = "S"
)

const (
	TypeMarket TypeT = "M"   // market: khớp theo bid/ask queue, không có giá. không được modify
	TypeLimit  TypeT = "L"   // đặt vào hàng chờ (modify)
	TypeATO    TypeT = "ATO" // đặt trước 9:00 (PS), trước 9g15 (CP), không giá, giá khớp là giá mở cửa. Không được modify
	TypeATC    TypeT = "ATC" // khớp 14:45, không được modify
)

const (
	StatusCreated         StatusT = "created"
	StatusPending         StatusT = "pending"
	StatusPartiallyFilled StatusT = "p_filled"
	StatusFullFilled      StatusT = "filled"

	// terminal states
	StatusCanceled StatusT = "canceled"
	StatusRejected StatusT = "rejected"
	StatusExpired  StatusT = "expired"
)

// Symbol types
type SymbolTypeT string

const (
	SymbolTypeVnStock  SymbolTypeT = "VnStock"
	SymbolTypeVnFuture SymbolTypeT = "VnFuture"
)
