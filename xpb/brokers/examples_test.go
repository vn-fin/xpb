package brokers

import (
	"context"
	"os"
	"testing"
)

//
//// TestGetOrderById - Lấy thông tin lệnh theo ID
//func TestGetOrderById(t *testing.T) {
//	if os.Getenv("BROKERS_GRPC_HOST") == "" {
//		os.Setenv("BROKERS_GRPC_HOST", "localhost:50052")
//	}
//
//	client, err := GetClient()
//	if err != nil {
//		t.Fatalf("Failed to create client: %v", err)
//	}
//	defer client.Close()
//
//	req := &GetOrderRequest{
//		AccessToken:    "your_access_token_here",
//		GroupId:        "5bec9e5e-d752-4340-845c-ac3c3590a076",
//		OrderSessionId: "f2020fee-e029-4317-a0e9-7d399b0782f4",
//	}
//
//	resp, err := client.GetOrderById(context.Background(), req)
//	if err != nil {
//		t.Fatalf("GetOrderById failed: %v", err)
//	}
//
//	t.Logf("GetOrderById Response:")
//	t.Logf("  Success: %v", resp.Success)
//	t.Logf("  Message: %s", resp.Message)
//	if resp.Order != nil {
//		t.Logf("  Order:")
//		t.Logf("    OrderSessionId: %s", resp.Order.OrderSessionId)
//		t.Logf("    Symbol: %s", resp.Order.Symbol)
//		t.Logf("    Price: %.2f", resp.Order.Price)
//		t.Logf("    Quantity: %.2f", resp.Order.Quantity)
//		t.Logf("    FilledQty: %.2f", resp.Order.FilledQty)
//		t.Logf("    Status: %s", resp.Order.Status)
//	}
//}

// TestGetPendingOrders - Lấy danh sách lệnh đang chờ
//func TestGetPendingOrders(t *testing.T) {
//	if os.Getenv("BROKERS_GRPC_HOST") == "" {
//		os.Setenv("BROKERS_GRPC_HOST", "localhost:50052")
//	}
//
//	client, err := GetClient()
//	if err != nil {
//		t.Fatalf("Failed to create client: %v", err)
//	}
//	defer client.Close()
//
//	req := &GetPendingOrdersRequest{
//		AccessToken: "your_access_token_here",
//		GroupId:     "5bec9e5e-d752-4340-845c-ac3c3590a076",
//	}
//
//	resp, err := client.GetPendingOrders(context.Background(), req)
//	if err != nil {
//		t.Fatalf("GetPendingOrders failed: %v", err)
//	}
//
//	t.Logf("GetPendingOrders Response:")
//	t.Logf("  Total pending orders: %d", len(resp.PendingOrders))
//	for i, order := range resp.PendingOrders {
//		t.Logf("  Order[%d]:", i)
//		t.Logf("    OrderSessionId: %s", order.OrderSessionId)
//		t.Logf("    Symbol: %s", order.Symbol)
//		t.Logf("    Quantity: %.2f", order.Quantity)
//		t.Logf("    Side: %s", order.Side)
//	}
//}

// TestCancelOrder - Hủy lệnh cổ phiếu
//func TestCancelOrder(t *testing.T) {
//	if os.Getenv("BROKERS_GRPC_HOST") == "" {
//		os.Setenv("BROKERS_GRPC_HOST", "localhost:50052")
//	}
//
//	client, err := GetClient()
//	if err != nil {
//		t.Fatalf("Failed to create client: %v", err)
//	}
//	defer client.Close()
//
//	req := &CancelOrderRequest{
//		AccessToken:    "your_access_token_here",
//		OrderSessionId: "84f48fd4-2dd4-49c9-91a8-14532fe4e9af",
//	}
//
//	resp, err := client.CancelOrder(context.Background(), req)
//	if err != nil {
//		t.Fatalf("CancelOrder failed: %v", err)
//	}
//
//	t.Logf("CancelOrder Response:")
//	t.Logf("  Success: %v", resp.Success)
//	t.Logf("  Message: %s", resp.Message)
//}
//
//// TestCancelFutureOrder - Hủy lệnh phái sinh
//func TestCancelFutureOrder(t *testing.T) {
//	if os.Getenv("BROKERS_GRPC_HOST") == "" {
//		os.Setenv("BROKERS_GRPC_HOST", "localhost:50052")
//	}
//
//	client, err := GetClient()
//	if err != nil {
//		t.Fatalf("Failed to create client: %v", err)
//	}
//	defer client.Close()
//
//	req := &CancelFutureOrderRequest{
//		AccessToken:    "your_access_token_here",
//		CredentialId:   "your_credential_id",
//		OrderSessionId: "your_order_session_id",
//	}
//
//	resp, err := client.CancelFutureOrder(context.Background(), req)
//	if err != nil {
//		t.Fatalf("CancelFutureOrder failed: %v", err)
//	}
//
//	t.Logf("CancelFutureOrder Response:")
//	t.Logf("  Success: %v", resp.Success)
//	t.Logf("  Message: %s", resp.Message)
//}

// TestGetPortfolio - Lấy danh mục đầu tư
//func TestGetPortfolio(t *testing.T) {
//	if os.Getenv("BROKERS_GRPC_HOST") == "" {
//		os.Setenv("BROKERS_GRPC_HOST", "localhost:50052")
//	}
//
//	client, err := GetClient()
//	if err != nil {
//		t.Fatalf("Failed to create client: %v", err)
//	}
//	defer client.Close()
//
//	req := &GetPortfolioRequest{
//		GroupId: "5bec9e5e-d752-4340-845c-ac3c3590a076",
//	}
//
//	resp, err := client.GetPortfolio(context.Background(), req)
//	if err != nil {
//		t.Fatalf("GetPortfolio failed: %v", err)
//	}
//
//	t.Logf("GetPortfolio Response:")
//	t.Logf("  Success: %v", resp.Success)
//	t.Logf("  Message: %s", resp.Message)
//	t.Logf("  Total items: %d", len(resp.Items))
//	for i, item := range resp.Items {
//		t.Logf("  Item[%d]:", i)
//		t.Logf("    Symbol: %s", item.Symbol)
//		t.Logf("    Quantity: %d", item.Quantity)
//		t.Logf("    AvgPrice: %.2f", item.AvgPrice)
//		t.Logf("    MarketPrice: %.2f", item.MarketPrice)
//	}
//}

// TestGetAccountBalance - Lấy số dư tài khoản
func TestGetAccountBalance(t *testing.T) {
	if os.Getenv("BROKERS_GRPC_HOST") == "" {
		os.Setenv("BROKERS_GRPC_HOST", "localhost:50052")
	}

	client, err := GetClient()
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	req := &GetAccountBalanceRequest{
		AccessToken: "5fbd07f06c212695250d0b0a925aa02082314018",
		GroupId:     "5bec9e5e-d752-4340-845c-ac3c3590a076",
		AccountType: "VnStock",
	}

	resp, err := client.GetAccountBalance(context.Background(), req)
	if err != nil {
		t.Fatalf("GetAccountBalance failed: %v", err)
	}

	t.Logf("GetAccountBalance Response:")
	t.Logf("  Success: %v", resp.Success)
	t.Logf("  Message: %s", resp.Message)
	t.Logf("  Cash: %.2f", resp.Cash)
	t.Logf("  Equity: %.2f", resp.Equity)
	t.Logf("  Balance: %.2f", resp.Balance)

}

// TestGetOrdersBySession - Lấy lệnh theo basket session
func TestGetOrdersBySession(t *testing.T) {
	if os.Getenv("BROKERS_GRPC_HOST") == "" {
		os.Setenv("BROKERS_GRPC_HOST", "localhost:50052")
	}

	client, err := GetClient()
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	req := &GetOrdersBySessionRequest{
		GroupId:         "5bec9e5e-d752-4340-845c-ac3c3590a076",
		BasketSessionId: "test-dnse-broker ",
	}

	resp, err := client.GetOrdersBySession(context.Background(), req)
	if err != nil {
		t.Fatalf("GetOrdersBySession failed: %v", err)
	}

	t.Logf("GetOrdersBySession Response:")
	t.Logf("  Success: %v", resp.Success)
	t.Logf("  Message: %s", resp.Message)
	t.Logf("  Total orders: %d", len(resp.Orders))
	for i, order := range resp.Orders {
		t.Logf("  Order[%d]:", i)
		t.Logf("    OrderSessionId: %s", order.OrderSessionId)
		t.Logf("    Symbol: %s", order.Symbol)
		t.Logf("    Status: %s", order.Status)
	}
}

// TestUpdateOrder - Cập nhật lệnh (hủy lệnh cũ, tạo lệnh mới với giá mới)
func TestUpdateOrder(t *testing.T) {
	if os.Getenv("BROKERS_GRPC_HOST") == "" {
		os.Setenv("BROKERS_GRPC_HOST", "100.74.98.123:50049")
	}

	client, err := GetClient()
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	req := &UpdateOrderRequest{
		AccessToken:    "your_access_token_here",
		GroupId:        "your_group_id",
		OrderSessionId: "your_order_session_id",
		Price:          85000, // New price
	}

	resp, err := client.UpdateOrder(context.Background(), req)
	if err != nil {
		t.Fatalf("UpdateOrder failed: %v", err)
	}

	t.Logf("UpdateOrder Response:")
	t.Logf("  Success: %v", resp.Success)
	t.Logf("  Message: %s", resp.Message)
	t.Logf("  OrderSessionId: %s", resp.OrderSessionId)
	if resp.Order != nil {
		t.Logf("  Order:")
		t.Logf("    OrderId: %s", resp.Order.OrderId)
		t.Logf("    Symbol: %s", resp.Order.Symbol)
		t.Logf("    Status: %s", resp.Order.Status)
	}
}

// TestCreateFutureOrder - Tạo lệnh phái sinh
func TestCreateFutureOrder(t *testing.T) {
	if os.Getenv("BROKERS_GRPC_HOST") == "" {
		os.Setenv("BROKERS_GRPC_HOST", "localhost:50052")
	}

	client, err := GetClient()
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	req := &CreateFutureOrderRequest{
		AccessToken: "your_access_token_here",
		UserId:      "your_user_id",
		Symbol:      "VN30F2412",
		OrderSide:   "B", // B (Buy) or S (Sell)
		OrderType:   "L", // L (Limit) or M (Market)
		Price:       1250,
		Quantity:    10,
	}

	resp, err := client.CreateFutureOrder(context.Background(), req)
	if err != nil {
		t.Fatalf("CreateFutureOrder failed: %v", err)
	}

	t.Logf("CreateFutureOrder Response:")
	t.Logf("  Success: %v", resp.Success)
	t.Logf("  Message: %s", resp.Message)
	if resp.Order != nil {
		t.Logf("  Order:")
		t.Logf("    Symbol: %s", resp.Order.Symbol)
		t.Logf("    OrderSessionId: %s", resp.Order.OrderSessionId)
	}
}

// TestGetListFutureOrders - Lấy danh sách lệnh phái sinh
func TestGetListFutureOrders(t *testing.T) {
	if os.Getenv("BROKERS_GRPC_HOST") == "" {
		os.Setenv("BROKERS_GRPC_HOST", "localhost:50052")
	}

	client, err := GetClient()
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	req := &GetListOrdersRequest{
		UserId: "your_user_id",
	}

	resp, err := client.GetListFutureOrders(context.Background(), req)
	if err != nil {
		t.Fatalf("GetListFutureOrders failed: %v", err)
	}

	t.Logf("GetListFutureOrders Response:")
	t.Logf("  Total orders: %d", len(resp.Orders))
	for i, order := range resp.Orders {
		t.Logf("  Order[%d]:", i)
		t.Logf("    OrderSessionId: %s", order.OrderSessionId)
		t.Logf("    Symbol: %s", order.Symbol)
		t.Logf("    Status: %s", order.Status)
	}
}

// TestGetPortfolioByGroupId - Lấy danh mục đầu tư theo group ID
func TestGetPortfolioByGroupId(t *testing.T) {
	if os.Getenv("BROKERS_GRPC_HOST") == "" {
		os.Setenv("BROKERS_GRPC_HOST", "localhost:50052")
	}

	client, err := GetClient()
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	req := &GetPortfolioRequest{
		GroupId: "your_group_id",
	}

	resp, err := client.GetPortfolioByGroupId(context.Background(), req)
	if err != nil {
		t.Fatalf("GetPortfolioByGroupId failed: %v", err)
	}

	t.Logf("GetPortfolioByGroupId Response:")
	t.Logf("  Success: %v", resp.Success)
	t.Logf("  Message: %s", resp.Message)
	t.Logf("  Total items: %d", len(resp.Items))
	for i, item := range resp.Items {
		t.Logf("  Item[%d]:", i)
		t.Logf("    Symbol: %s", item.Symbol)
		t.Logf("    Quantity: %d", item.Quantity)
	}
}

// TestExecutionGetPositions - Lấy tất cả vị thế từ execution service
func TestExecutionGetPositions(t *testing.T) {
	if os.Getenv("BROKERS_GRPC_HOST") == "" {
		os.Setenv("BROKERS_GRPC_HOST", "localhost:50052")
	}

	client, err := GetClient()
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	req := &ExecutionGetPositionsRequest{
		AccessToken: "your_access_token_here",
		CredId:      "asdd122e-asdx-adin-xzca-12313asdccsa",
	}

	resp, err := client.ExecutionGetPositions(context.Background(), req)
	if err != nil {
		t.Fatalf("ExecutionGetPositions failed: %v", err)
	}

	t.Logf("ExecutionGetPositions Response:")
	t.Logf("  Total positions: %d", len(resp.Positions))
	for symbol, pos := range resp.Positions {
		t.Logf("  Position[%s]:", symbol)
		t.Logf("    Symbol: %s", pos.Symbol)
		t.Logf("    Position: %.2f", pos.Position)
		if pos.Sellable != nil {
			t.Logf("    Sellable: %.2f", *pos.Sellable)
		}
	}
}

// TestExecutionGetPositionBySymbol - Lấy vị thế theo mã CK từ execution service
func TestExecutionGetPositionBySymbol(t *testing.T) {
	if os.Getenv("BROKERS_GRPC_HOST") == "" {
		os.Setenv("BROKERS_GRPC_HOST", "localhost:50052")
	}

	client, err := GetClient()
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	req := &ExecutionGetPositionBySymbolRequest{
		AccessToken: "your_access_token_here",
		CredId:      "asdd122e-asdx-adin-xzca-12313asdccsa",
		Symbol:      "VNM",
	}

	resp, err := client.ExecutionGetPositionBySymbol(context.Background(), req)
	if err != nil {
		t.Fatalf("ExecutionGetPositionBySymbol failed: %v", err)
	}

	t.Logf("ExecutionGetPositionBySymbol Response:")
	t.Logf("  Position: %.2f", resp.Position)
	if resp.Sellable != nil {
		t.Logf("  Sellable: %.2f", *resp.Sellable)
	}
}

// TestExecutionCancelOrder - Hủy lệnh từ execution service
func TestExecutionCancelOrder(t *testing.T) {
	if os.Getenv("BROKERS_GRPC_HOST") == "" {
		os.Setenv("BROKERS_GRPC_HOST", "localhost:50052")
	}

	client, err := GetClient()
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	req := &ExecutionCancelOrderRequest{
		AccessToken:    "your_access_token_here",
		CredId:         "asdd122e-asdx-adin-xzca-12313asdccsa",
		OrderSessionId: "your_order_session_id",
	}

	resp, err := client.ExecutionCancelOrder(context.Background(), req)
	if err != nil {
		t.Fatalf("ExecutionCancelOrder failed: %v", err)
	}

	t.Logf("ExecutionCancelOrder Response:")
	t.Logf("  Success: %v", resp.Success)
	t.Logf("  Message: %s", resp.Message)
}

// TestExecutionGetPendingOrders - Lấy danh sách lệnh chờ từ execution service
func TestExecutionGetPendingOrders(t *testing.T) {
	if os.Getenv("BROKERS_GRPC_HOST") == "" {
		os.Setenv("BROKERS_GRPC_HOST", "localhost:50052")
	}

	client, err := GetClient()
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	req := &ExecutionGetPendingOrdersRequest{
		AccessToken: "your_access_token_here",
		CredId:      "asdd122e-asdx-adin-xzca-12313asdccsa",
	}

	resp, err := client.ExecutionGetPendingOrders(context.Background(), req)
	if err != nil {
		t.Fatalf("ExecutionGetPendingOrders failed: %v", err)
	}

	t.Logf("ExecutionGetPendingOrders Response:")
	t.Logf("  Total pending orders: %d", len(resp.PendingOrders))
	for i, order := range resp.PendingOrders {
		t.Logf("  Order[%d]:", i)
		t.Logf("    OrderSessionId: %s", order.OrderSessionId)
		t.Logf("    Symbol: %s", order.Symbol)
		t.Logf("    Quantity: %.2f", order.Quantity)
	}
}

// TestExecutionGetPendingOrdersBySymbol - Lấy lệnh chờ theo mã CK từ execution service
func TestExecutionGetPendingOrdersBySymbol(t *testing.T) {
	if os.Getenv("BROKERS_GRPC_HOST") == "" {
		os.Setenv("BROKERS_GRPC_HOST", "localhost:50052")
	}

	client, err := GetClient()
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	req := &ExecutionGetPendingOrdersBySymbolRequest{
		AccessToken: "your_access_token_here",
		CredId:      "asdd122e-asdx-adin-xzca-12313asdccsa",
		Symbol:      "VNM",
	}

	resp, err := client.ExecutionGetPendingOrdersBySymbol(context.Background(), req)
	if err != nil {
		t.Fatalf("ExecutionGetPendingOrdersBySymbol failed: %v", err)
	}

	t.Logf("ExecutionGetPendingOrdersBySymbol Response:")
	t.Logf("  Total pending orders: %d", len(resp.PendingOrders))
	for i, order := range resp.PendingOrders {
		t.Logf("  Order[%d]:", i)
		t.Logf("    OrderSessionId: %s", order.OrderSessionId)
		t.Logf("    Symbol: %s", order.Symbol)
		t.Logf("    Quantity: %.2f", order.Quantity)
	}
}

// TestExecutionGetMaxSellQtyBySymbol - Lấy số lượng bán tối đa theo mã CK
func TestExecutionGetMaxSellQtyBySymbol(t *testing.T) {
	if os.Getenv("BROKERS_GRPC_HOST") == "" {
		os.Setenv("BROKERS_GRPC_HOST", "localhost:50052")
	}

	client, err := GetClient()
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	req := &ExecutionGetMaxSellQtyBySymbolRequest{
		AccessToken: "your_access_token_here",
		CredId:      "asdd122e-asdx-adin-xzca-12313asdccsa",
		Symbol:      "VNM",
	}

	resp, err := client.ExecutionGetMaxSellQtyBySymbol(context.Background(), req)
	if err != nil {
		t.Fatalf("ExecutionGetMaxSellQtyBySymbol failed: %v", err)
	}

	t.Logf("ExecutionGetMaxSellQtyBySymbol Response:")
	t.Logf("  MaxSellQty: %.2f", resp.MaxSellQty)
}

// TestExecutionGetMaxSellQtys - Lấy tất cả số lượng bán tối đa
func TestExecutionGetMaxSellQtys(t *testing.T) {
	if os.Getenv("BROKERS_GRPC_HOST") == "" {
		os.Setenv("BROKERS_GRPC_HOST", "localhost:50052")
	}

	client, err := GetClient()
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	req := &ExecutionGetMaxSellQtysRequest{
		AccessToken: "your_access_token_here",
		CredId:      "asdd122e-asdx-adin-xzca-12313asdccsa",
	}

	resp, err := client.ExecutionGetMaxSellQtys(context.Background(), req)
	if err != nil {
		t.Fatalf("ExecutionGetMaxSellQtys failed: %v", err)
	}

	t.Logf("ExecutionGetMaxSellQtys Response:")
	t.Logf("  Total symbols: %d", len(resp.MaxSellInfos))
	for symbol, info := range resp.MaxSellInfos {
		t.Logf("  Symbol[%s]:", symbol)
		t.Logf("    Symbol: %s", info.Symbol)
		if info.MaxSellQtys != nil {
			t.Logf("    MaxSellQty: %.2f", *info.MaxSellQtys)
		}
	}
}
