package trading212

import (
	"context"
	"fmt"
	"net/http"
)

// MarketOrderRequest represents a market order request
type MarketOrderRequest struct {
	ExtendedHours bool    `json:"extendedHours"`
	Quantity      float64 `json:"quantity"`
	Ticker        string  `json:"ticker"`
}

// LimitOrderRequest represents a limit order request
type LimitOrderRequest struct {
	LimitPrice   float64      `json:"limitPrice"`
	Quantity     float64      `json:"quantity"`
	Ticker       string       `json:"ticker"`
	TimeValidity TimeValidity `json:"timeValidity"`
}

// StopOrderRequest represents a stop order request
type StopOrderRequest struct {
	Quantity     float64      `json:"quantity"`
	StopPrice    float64      `json:"stopPrice"`
	Ticker       string       `json:"ticker"`
	TimeValidity TimeValidity `json:"timeValidity"`
}

// StopLimitOrderRequest represents a stop-limit order request
type StopLimitOrderRequest struct {
	LimitPrice   float64      `json:"limitPrice"`
	Quantity     float64      `json:"quantity"`
	StopPrice    float64      `json:"stopPrice"`
	Ticker       string       `json:"ticker"`
	TimeValidity TimeValidity `json:"timeValidity"`
}

// GetOrders retrieves all pending orders
func (c *Client) GetOrders(ctx context.Context) ([]Order, error) {
	resp, err := c.makeRequest(ctx, http.MethodGet, "/api/v0/equity/orders", nil)
	if err != nil {
		return nil, err
	}

	var orders []Order
	if err := c.handleResponse(resp, &orders); err != nil {
		return nil, err
	}

	return orders, nil
}

// GetOrderByID retrieves a specific order by ID
func (c *Client) GetOrderByID(ctx context.Context, orderID int64) (*Order, error) {
	path := fmt.Sprintf("/api/v0/equity/orders/%d", orderID)
	resp, err := c.makeRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var order Order
	if err := c.handleResponse(resp, &order); err != nil {
		return nil, err
	}

	return &order, nil
}

// PlaceMarketOrder places a market order
func (c *Client) PlaceMarketOrder(ctx context.Context, req MarketOrderRequest) (*Order, error) {
	resp, err := c.makeRequest(ctx, http.MethodPost, "/api/v0/equity/orders/market", req)
	if err != nil {
		return nil, err
	}

	var order Order
	if err := c.handleResponse(resp, &order); err != nil {
		return nil, err
	}

	return &order, nil
}

// PlaceLimitOrder places a limit order
func (c *Client) PlaceLimitOrder(ctx context.Context, req LimitOrderRequest) (*Order, error) {
	resp, err := c.makeRequest(ctx, http.MethodPost, "/api/v0/equity/orders/limit", req)
	if err != nil {
		return nil, err
	}

	var order Order
	if err := c.handleResponse(resp, &order); err != nil {
		return nil, err
	}

	return &order, nil
}

// PlaceStopOrder places a stop order
func (c *Client) PlaceStopOrder(ctx context.Context, req StopOrderRequest) (*Order, error) {
	resp, err := c.makeRequest(ctx, http.MethodPost, "/api/v0/equity/orders/stop", req)
	if err != nil {
		return nil, err
	}

	var order Order
	if err := c.handleResponse(resp, &order); err != nil {
		return nil, err
	}

	return &order, nil
}

// PlaceStopLimitOrder places a stop-limit order
func (c *Client) PlaceStopLimitOrder(ctx context.Context, req StopLimitOrderRequest) (*Order, error) {
	resp, err := c.makeRequest(ctx, http.MethodPost, "/api/v0/equity/orders/stop_limit", req)
	if err != nil {
		return nil, err
	}

	var order Order
	if err := c.handleResponse(resp, &order); err != nil {
		return nil, err
	}

	return &order, nil
}

// CancelOrder cancels an order by ID
func (c *Client) CancelOrder(ctx context.Context, orderID int64) error {
	path := fmt.Sprintf("/api/v0/equity/orders/%d", orderID)
	resp, err := c.makeRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return err
	}

	return c.handleResponse(resp, nil)
}