package brokers

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

const (
	AuthorizationHeader = "authorization"
)

// Client wraps the BrokerGatewayServiceClient with authorization support
type Client struct {
	conn   *grpc.ClientConn
	client BrokerGatewayServiceClient
	token  string
}

// ClientOption is a function that configures a Client
type ClientOption func(*Client)

// WithToken sets the authorization token
func WithToken(token string) ClientOption {
	return func(c *Client) {
		c.token = token
	}
}

// NewClient creates a new broker client with the given server address
func NewClient(serverAddr string, opts ...ClientOption) (*Client, error) {
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

	for _, opt := range opts {
		opt(c)
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

// SetToken sets or updates the authorization token
func (c *Client) SetToken(token string) {
	c.token = token
}

// contextWithAuth returns a context with authorization metadata
func (c *Client) contextWithAuth(ctx context.Context) context.Context {
	if c.token == "" {
		return ctx
	}
	return metadata.AppendToOutgoingContext(ctx, AuthorizationHeader, c.token)
}

// Login authenticates with the broker
func (c *Client) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	return c.client.Login(c.contextWithAuth(ctx), req)
}

// CreateOrder creates a new order
func (c *Client) CreateOrder(ctx context.Context, req *CreateOrderRequest) (*CreateOrderResponse, error) {
	return c.client.CreateOrder(c.contextWithAuth(ctx), req)
}

// GetOrderById retrieves an order by its ID
func (c *Client) GetOrderById(ctx context.Context, req *GetOrderByIdRequest) (*GetOrderByIdResponse, error) {
	return c.client.GetOrderById(c.contextWithAuth(ctx), req)
}

// GetPendingOrders retrieves all pending orders
func (c *Client) GetPendingOrders(ctx context.Context, req *GetPendingOrdersRequest) (*GetPendingOrdersResponse, error) {
	return c.client.GetPendingOrders(c.contextWithAuth(ctx), req)
}

// CancelOrder cancels an order
func (c *Client) CancelOrder(ctx context.Context, req *CancelOrderRequest) (*CancelOrderResponse, error) {
	return c.client.CancelOrder(c.contextWithAuth(ctx), req)
}

// GetPortfolio retrieves the portfolio
func (c *Client) GetPortfolio(ctx context.Context, req *GetPortfolioRequest) (*GetPortfolioResponse, error) {
	return c.client.GetPortfolio(c.contextWithAuth(ctx), req)
}

// GetAccountBalance retrieves the account balance
func (c *Client) GetAccountBalance(ctx context.Context, req *GetAccountBalanceRequest) (*GetAccountBalanceResponse, error) {
	return c.client.GetAccountBalance(c.contextWithAuth(ctx), req)
}
