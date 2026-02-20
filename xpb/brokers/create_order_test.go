package brokers

import (
	"context"
	"os"
	"testing"
)

func TestCreateOrder(t *testing.T) {
	if os.Getenv("BROKERS_GRPC_HOST") == "" {
		os.Setenv("BROKERS_GRPC_HOST", "localhost:50052")
	}

	client, err := GetClient()
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	req := &CreateOrderRequest{
		AccessToken:     "5fbd07f06c212695250d0b0a925aa02082314018", // Replace with actual access token from Login
		GroupId:         "5bec9e5e-d752-4340-845c-ac3c3590a076",
		BasketSessionId: "test-dnse-broker ",
		OrderInfo: []*OrderInputInfo{
			{
				Symbol:     "OIL",
				SymbolType: "VnStock",
				Price:      15.30,
				Quantity:   1,
				OrderType:  "L", // LO, MP, ATO, ATC
				OrderSide:  "B", // buy, sell
			},
			{
				Symbol:     "PVC",
				SymbolType: "VnStock",
				Price:      14.30,
				Quantity:   1,
				OrderType:  "L", // LO, MP, ATO, ATC
				OrderSide:  "B", // buy, sell
			},
		},
	}

	resp, err := client.CreateOrder(context.Background(), req)
	if err != nil {
		t.Fatalf("CreateOrder failed: %v", err)
	}

	t.Logf("CreateOrder Response:")
	t.Logf("  Success: %v", resp.Success)
	t.Logf("  Message: %s", resp.Message)
	for i, order := range resp.Orders {
		t.Logf("  Order[%d]: Symbol=%s, OrderSessionId=%s", i, order.Symbol, order.OrderSessionId)
	}

	if !resp.Success {
		t.Errorf("Expected success=true, got success=%v, message=%s", resp.Success, resp.Message)
	}
}
