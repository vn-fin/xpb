package brokers

import (
	"context"
	"os"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	defaultClient *Client
	clientMu      sync.Mutex
)

// Client wraps the BrokerGatewayServiceClient
type Client struct {
	conn   *grpc.ClientConn
	client BrokerGatewayServiceClient
}

// GetClient returns a singleton client instance.
// It initializes the client if it doesn't exist.
func GetClient() (*Client, error) {
	clientMu.Lock()
	defer clientMu.Unlock()

	if defaultClient != nil {
		return defaultClient, nil
	}

	c, err := NewClient()
	if err != nil {
		return nil, err
	}
	defaultClient = c
	return defaultClient, nil
}

// NewClient creates a new broker client with the given server address

func NewClient() (*Client, error) {
	serverAddr := os.Getenv("BROKERS_GRPC_HOST")
	conn, err := grpc.NewClient(serverAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	c := &Client{
		conn:   conn,
		client: NewBrokerGatewayServiceClient(conn),
	}

	return c, nil
}

// Close closes the connection
func (c *Client) Close() error {
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}

// Login authenticates with the broker
func (c *Client) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	return c.client.Login(ctx, req)
}

// CreateOrder creates a new order
func (c *Client) CreateOrder(ctx context.Context, req *CreateOrderRequest) (*CreateOrderResponse, error) {
	return c.client.CreateOrder(ctx, req)
}

// GetOrderById retrieves an order by its ID
func (c *Client) GetOrderById(ctx context.Context, req *GetOrderRequest) (*GetOrderResponse, error) {
	return c.client.GetOrderById(ctx, req)
}

// GetPendingOrders retrieves all pending orders
func (c *Client) GetPendingOrders(ctx context.Context, req *GetPendingOrdersRequest) (*GetPendingOrdersResponse, error) {
	return c.client.GetPendingOrders(ctx, req)
}

// CancelOrder cancels an order
func (c *Client) CancelOrder(ctx context.Context, req *CancelOrderRequest) (*CancelOrderResponse, error) {
	return c.client.CancelOrder(ctx, req)
}

// CancelFutureOrder cancels an order
func (c *Client) CancelFutureOrder(ctx context.Context, req *CancelFutureOrderRequest) (*CancelOrderResponse, error) {
	return c.client.CancelFutureOrder(ctx, req)
}

// GetPortfolio retrieves the portfolio
func (c *Client) GetPortfolio(ctx context.Context, req *GetPortfolioRequest) (*GetPortfolioResponse, error) {
	return c.client.GetPortfolio(ctx, req)
}

// GetAccountBalance retrieves the account balance
func (c *Client) GetAccountBalance(ctx context.Context, req *GetAccountBalanceRequest) (*GetAccountBalanceResponse, error) {
	return c.client.GetAccountBalance(ctx, req)
}

func (c *Client) GetOrdersBySession(ctx context.Context, req *GetOrdersBySessionRequest) (*GetOrdersBySessionResponse, error) {
	return c.client.GetOrdersBySession(ctx, req)
}

// UpdateOrder retrieves the account balance
func (c *Client) UpdateOrder(ctx context.Context, req *UpdateOrderRequest) (*UpdateOrderResponse, error) {
	return c.client.UpdateOrder(ctx, req)
}

// CreateFutureOrder retrieves the account balance
func (c *Client) CreateFutureOrder(ctx context.Context, req *CreateFutureOrderRequest) (*CreateFutureOrderResponse, error) {
	return c.client.CreateFutureOrder(ctx, req)
}

func (c *Client) GetListFutureOrders(ctx context.Context, req *GetListOrdersRequest) (*GetListOrdersResponse, error) {
	return c.client.GetListFutureOrders(ctx, req)
}
