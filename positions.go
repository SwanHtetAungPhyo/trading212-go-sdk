package trading212

import (
	"context"
	"net/http"
)

// GetPositionsOptions represents options for getting positions
type GetPositionsOptions struct {
	Ticker string
}

// GetPositions retrieves all open positions
func (c *Client) GetPositions(ctx context.Context, opts *GetPositionsOptions) ([]Position, error) {
	path := "/api/v0/equity/portfolio"
	
	if opts != nil {
		params := map[string]interface{}{
			"ticker": opts.Ticker,
		}
		path += buildQuery(params)
	}

	resp, err := c.makeRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var positions []Position
	if err := c.handleResponse(resp, &positions); err != nil {
		return nil, err
	}

	return positions, nil
}