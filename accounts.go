package trading212

import (
	"context"
	"net/http"
)

// GetAccountInfo retrieves account metadata information (currency, ID)
func (c *Client) GetAccountInfo(ctx context.Context) (*AccountInfo, error) {
	resp, err := c.makeRequest(ctx, http.MethodGet, "/api/v0/equity/account/info", nil)
	if err != nil {
		return nil, err
	}

	var info AccountInfo
	if err := c.handleResponse(resp, &info); err != nil {
		return nil, err
	}

	return &info, nil
}

// GetAccountCash retrieves account cash balance information
func (c *Client) GetAccountCash(ctx context.Context) (*AccountCash, error) {
	resp, err := c.makeRequest(ctx, http.MethodGet, "/api/v0/equity/account/cash", nil)
	if err != nil {
		return nil, err
	}

	var cash AccountCash
	if err := c.handleResponse(resp, &cash); err != nil {
		return nil, err
	}

	return &cash, nil
}

// GetAccountSummary retrieves combined account information (convenience method)
func (c *Client) GetAccountSummary(ctx context.Context) (*AccountSummary, error) {
	// Get account info
	info, err := c.GetAccountInfo(ctx)
	if err != nil {
		return nil, err
	}

	// Get cash balance
	cash, err := c.GetAccountCash(ctx)
	if err != nil {
		return nil, err
	}

	// Combine into summary
	summary := &AccountSummary{
		AccountInfo: *info,
		Cash:        *cash,
	}

	return summary, nil
}