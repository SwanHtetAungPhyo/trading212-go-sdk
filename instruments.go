package trading212

import (
	"context"
	"net/http"
)

// GetInstruments retrieves all available instruments
func (c *Client) GetInstruments(ctx context.Context) ([]TradableInstrument, error) {
	resp, err := c.makeRequest(ctx, http.MethodGet, "/api/v0/equity/metadata/instruments", nil)
	if err != nil {
		return nil, err
	}

	var instruments []TradableInstrument
	if err := c.handleResponse(resp, &instruments); err != nil {
		return nil, err
	}

	return instruments, nil
}

// GetExchanges retrieves all accessible exchanges
func (c *Client) GetExchanges(ctx context.Context) ([]Exchange, error) {
	resp, err := c.makeRequest(ctx, http.MethodGet, "/api/v0/equity/metadata/exchanges", nil)
	if err != nil {
		return nil, err
	}

	var exchanges []Exchange
	if err := c.handleResponse(resp, &exchanges); err != nil {
		return nil, err
	}

	return exchanges, nil
}