package brokers

import (
	"context"
	"os"
	"testing"
)

func TestExecutionCreateOrder(t *testing.T) {
	if os.Getenv("BROKERS_GRPC_HOST") == "" {
		os.Setenv("BROKERS_GRPC_HOST", "localhost:50052")
	}

	client, err := GetClient()
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	req := &ExecutionCreateOrderRequest{
		AccessToken: "your_access_token_here",
		CredId:      "asdd122e-asdx-adin-xzca-12313asdccsa",
		Symbol:      "OIL",
		OrderSide:   "B",
		OrderType:   "L",
		Price:       16,
		Quantity:    1,
	}

	resp, err := client.ExecutionCreateOrder(context.Background(), req)
	if err != nil {
		t.Fatalf("ExecutionCreateOrder failed: %v", err)
	}

	t.Logf("ExecutionCreateOrder Response:")
	t.Logf("  OriginalOrderId: %s", resp.OriginalOrderId)
	t.Logf("  OrderSessionId: %s", resp.OrderSessionId)
	for i, order := range resp.Orders {
		t.Logf("  Order[%d]: Symbol=%s, OrderSessionId=%s", i, order.Symbol, order.OrderSessionId)
	}
}
