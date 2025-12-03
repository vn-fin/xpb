package order

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
	TypeMarket TypeT = "M"   // market: khớp theo bid/ask queue, không có giá. không được modify
	TypeLimit  TypeT = "L"   // đặt vào hàng chờ (modify)
	TypeATO    TypeT = "ATO" // đặt trước 9:00 (PS), trước 9g15 (CP), không giá, giá khớp là giá mở cửa. Không được modify
	TypeATC    TypeT = "ATC" // khớp 14:45, không được modify
)

const (
	StatusCreated         StatusT = 1 // vừa đặt lệnh
	StatusPending         StatusT = 2 // đang chờ khớp
	StatusFilled          StatusT = 3 // khớp hoàn toàn
	StatusCanceled        StatusT = 4 // đã hủy
	StatusRejected        StatusT = 5 // bị reject (random rate 1%), ngoài giờ giao dịch
	StatusExpired         StatusT = 6 // sau 3g15, nếu chưa khớp -> hết hạn (cron)
	StatusPartiallyFilled StatusT = 7 // khớp 1 phần: mua 10k, khớp 500
)
