package trading212

import (
	"context"
	"net/http"
)

// GetAccountSummary retrieves account summary information
func (c *Client) GetAccountSummary(ctx context.Context) (*AccountSummary, error) {
	resp, err := c.makeRequest(ctx, http.MethodGet, "/api/v0/equity/account/summary", nil)
	if err != nil {
		return nil, err
	}

	var summary AccountSummary
	if err := c.handleResponse(resp, &summary); err != nil {
		return nil, err
	}

	return &summary, nil
}