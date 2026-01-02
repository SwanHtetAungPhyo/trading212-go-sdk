package trading212

import (
	"context"
	"net/http"
	"time"
)

// HistoricalOrder represents a historical order
type HistoricalOrder struct {
	Fill  Fill  `json:"fill"`
	Order Order `json:"order"`
}

// Fill represents order fill information
type Fill struct {
	FilledAt      time.Time        `json:"filledAt"`
	ID            int64            `json:"id"`
	Price         float64          `json:"price"`
	Quantity      float64          `json:"quantity"`
	TradingMethod string           `json:"tradingMethod"`
	Type          string           `json:"type"`
	WalletImpact  FillWalletImpact `json:"walletImpact"`
}

// FillWalletImpact represents fill wallet impact
type FillWalletImpact struct {
	Currency           string  `json:"currency"`
	FxRate             float64 `json:"fxRate"`
	NetValue           float64 `json:"netValue"`
	RealisedProfitLoss float64 `json:"realisedProfitLoss"`
	Taxes              []Tax   `json:"taxes"`
}

// Tax represents tax information
type Tax struct {
	ChargedAt time.Time `json:"chargedAt"`
	Currency  string    `json:"currency"`
	Name      string    `json:"name"`
	Quantity  float64   `json:"quantity"`
}

// HistoryDividendItem represents a dividend item
type HistoryDividendItem struct {
	Amount               float64    `json:"amount"`
	AmountInEuro         float64    `json:"amountInEuro"`
	Currency             string     `json:"currency"`
	GrossAmountPerShare  float64    `json:"grossAmountPerShare"`
	Instrument           Instrument `json:"instrument"`
	PaidOn               time.Time  `json:"paidOn"`
	Quantity             float64    `json:"quantity"`
	Reference            string     `json:"reference"`
	Ticker               string     `json:"ticker"`
	TickerCurrency       string     `json:"tickerCurrency"`
	Type                 string     `json:"type"`
}

// HistoryTransactionItem represents a transaction item
type HistoryTransactionItem struct {
	Amount    float64   `json:"amount"`
	Currency  string    `json:"currency"`
	DateTime  time.Time `json:"dateTime"`
	Reference string    `json:"reference"`
	Type      string    `json:"type"`
}

// PaginatedResponse represents a paginated response
type PaginatedResponse[T any] struct {
	Items        []T     `json:"items"`
	NextPagePath *string `json:"nextPagePath"`
}

// HistoryOrdersOptions represents options for getting historical orders
type HistoryOrdersOptions struct {
	Cursor int64
	Ticker string
	Limit  int
}

// HistoryDividendsOptions represents options for getting dividends
type HistoryDividendsOptions struct {
	Cursor int64
	Ticker string
	Limit  int
}

// HistoryTransactionsOptions represents options for getting transactions
type HistoryTransactionsOptions struct {
	Cursor string
	Time   *time.Time
	Limit  int
}

// GetHistoricalOrders retrieves historical orders
func (c *Client) GetHistoricalOrders(ctx context.Context, opts *HistoryOrdersOptions) (*PaginatedResponse[HistoricalOrder], error) {
	path := "/api/v0/equity/history/orders"
	
	if opts != nil {
		params := map[string]interface{}{
			"cursor": opts.Cursor,
			"ticker": opts.Ticker,
			"limit":  opts.Limit,
		}
		path += buildQuery(params)
	}

	resp, err := c.makeRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var result PaginatedResponse[HistoricalOrder]
	if err := c.handleResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetDividends retrieves dividend history
func (c *Client) GetDividends(ctx context.Context, opts *HistoryDividendsOptions) (*PaginatedResponse[HistoryDividendItem], error) {
	path := "/api/v0/equity/history/dividends"
	
	if opts != nil {
		params := map[string]interface{}{
			"cursor": opts.Cursor,
			"ticker": opts.Ticker,
			"limit":  opts.Limit,
		}
		path += buildQuery(params)
	}

	resp, err := c.makeRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var result PaginatedResponse[HistoryDividendItem]
	if err := c.handleResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetTransactions retrieves transaction history
func (c *Client) GetTransactions(ctx context.Context, opts *HistoryTransactionsOptions) (*PaginatedResponse[HistoryTransactionItem], error) {
	path := "/api/v0/equity/history/transactions"
	
	if opts != nil {
		params := map[string]interface{}{
			"cursor": opts.Cursor,
			"time":   opts.Time,
			"limit":  opts.Limit,
		}
		path += buildQuery(params)
	}

	resp, err := c.makeRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var result PaginatedResponse[HistoryTransactionItem]
	if err := c.handleResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}